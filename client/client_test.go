package main

import (
	"context"
	"errors"
	"grpc/mock"
	"grpc/protobuf/pb"
	"io"
	"log"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestClient_SearchPoint(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRendererClient := mock.NewMockRendererClient(ctrl)

	mockRendererClient.EXPECT().SearchPoint(gomock.Any(), &pb.Point{X: 10, Y: 2}).Return(&pb.Triangle{TrId: 1, Vertice_1: &pb.Point{}, Vertice_2: &pb.Point{}, Vertice_3: &pb.Point{}}, nil)
	mockRendererClient.EXPECT().SearchPoint(gomock.Any(), &pb.Point{}).Return(nil, errors.New("No Triangle with a Point as one of its vertices"))

	triangle, err := mockRendererClient.SearchPoint(context.Background(), &pb.Point{X: 10, Y: 2})
	if err != nil {
		log.Fatalf("Error occoured %v", err)
	}
	log.Printf("Triangle received %+01v", triangle)

	triangle, err = mockRendererClient.SearchPoint(context.Background(), &pb.Point{})
	if err == nil {
		log.Fatalf("Error occoured %v", err)
	}
	log.Printf("Triangle received %+01v", triangle)

}

func TestClient_CreateTriangle(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRendererClient := mock.NewMockRendererClient(ctrl)
	mockRenderer_CreateTriangleClient := mock.NewMockRenderer_CreateTriangleClient(ctrl)

	testTriangleVerticePoints := []*pb.Point{
		{X: 1, Y: 3},
		{X: 5, Y: 4},
		{X: 9, Y: 12},
	}

	mockRendererClient.EXPECT().CreateTriangle(gomock.Any()).Return(mockRenderer_CreateTriangleClient, nil)
	for _, verticePoint := range testTriangleVerticePoints {
		mockRenderer_CreateTriangleClient.EXPECT().Send(verticePoint).Return(nil)
	}
	mockRenderer_CreateTriangleClient.EXPECT().CloseAndRecv().Return(&pb.Triangle{TrId: 1, Vertice_1: testTriangleVerticePoints[0], Vertice_2: testTriangleVerticePoints[1], Vertice_3: testTriangleVerticePoints[2]}, nil)

	mockRendererClient.EXPECT().CreateTriangle(gomock.Any()).Return(mockRenderer_CreateTriangleClient, nil)
	mockRenderer_CreateTriangleClient.EXPECT().CloseAndRecv().Return(nil, errors.New("err:in the receive stream"))

	_, err := executeCreateTriangle(testTriangleVerticePoints, mockRendererClient)
	if err != nil {
		log.Fatalf("Error occoured %v", err)
	}

	_, err = executeCreateTriangle(nil, mockRendererClient)
	if err == nil {
		log.Fatalf("Error occoured %v", err)
	}
}

func TestClient_GetPolyTriangles(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRendererClient := mock.NewMockRendererClient(ctrl)
	mockRenderer_GetPolyTrianglesClient := mock.NewMockRenderer_GetPolyTrianglesClient(ctrl)

	testTrianglesSplice := []*pb.Triangle{
		{TrId: 1, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 2, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 3, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
	}

	testPolygonTrianglesMap := []struct {
		polygon   *pb.Polygon
		triangles []*pb.Triangle
		err       error
		callErr   error
	}{
		{polygon: &pb.Polygon{}, triangles: nil, err: errors.New("Empty Polygon"), callErr: errors.New("No transaction on stream")},
		{polygon: &pb.Polygon{}, triangles: nil, err: errors.New("Empty Polygon"), callErr: nil},
		{polygon: &pb.Polygon{PolyName: "First Polygon", NumTriangles: 3, Triangles: testTrianglesSplice}, triangles: testTrianglesSplice, err: nil, callErr: nil},
	}

	for _, struct_tmp := range testPolygonTrianglesMap {
		mockRendererClient.EXPECT().GetPolyTriangles(gomock.Any(), struct_tmp.polygon).Return(mockRenderer_GetPolyTrianglesClient, struct_tmp.callErr)
		for _, triangle := range struct_tmp.triangles {
			mockRenderer_GetPolyTrianglesClient.EXPECT().Recv().Return(triangle, nil)
		}

		if struct_tmp.triangles == nil && struct_tmp.callErr == nil {
			mockRenderer_GetPolyTrianglesClient.EXPECT().Recv().Return(nil, struct_tmp.err)
		} else if struct_tmp.callErr == nil {
			mockRenderer_GetPolyTrianglesClient.EXPECT().Recv().Return(nil, io.EOF)
		}
		err := executeGetPolyTriangles(mockRendererClient, struct_tmp.polygon)
		if err != nil {
			log.Printf("Error fatal: %v", err)
		}
	}
}

func TestClient_GetTriangleList(t *testing.T) {
	ctrl := gomock.NewController(t)

	testTrianglesSplice := []*pb.Triangle{
		{TrId: 1, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 2, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 3, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
	}

	testTriangleList := []struct {
		triangles []*pb.Triangle
		err       error
		callErr   error
	}{
		{triangles: testTrianglesSplice, err: nil, callErr: nil},
		{triangles: nil, err: errors.New("Recv"), callErr: nil},
		{triangles: nil, err: nil, callErr: errors.New("GetTriangleList")},
	}

	mockRendererClient := mock.NewMockRendererClient(ctrl)
	mockRenderer_ListOfTrianglesClient := mock.NewMockRenderer_ListOfTrianglesClient(ctrl)

	for _, testStruct := range testTriangleList {
		mockRendererClient.EXPECT().ListOfTriangles(gomock.Any(), nil).Return(mockRenderer_ListOfTrianglesClient, testStruct.callErr)
		for _, triangle := range testStruct.triangles {
			mockRenderer_ListOfTrianglesClient.EXPECT().Recv().Return(triangle, nil)
		}
		if testStruct.triangles == nil && testStruct.callErr == nil {
			mockRenderer_ListOfTrianglesClient.EXPECT().Recv().Return(nil, testStruct.err)
		} else if testStruct.callErr == nil {
			mockRenderer_ListOfTrianglesClient.EXPECT().Recv().Return(nil, io.EOF)
		}

		_, err := getTriangleList(mockRendererClient)
		if err != nil {
			log.Printf("Error caused by %v", err)
		}

	}

}

func TestClient_GetPolygonList(t *testing.T) {
	ctrl := gomock.NewController(t)

	testPolygonSplice := []*pb.Polygon{
		{PolyName: "Polygon 1", NumTriangles: 3, Triangles: []*pb.Triangle{}},
		{PolyName: "Polygon 2", NumTriangles: 2, Triangles: []*pb.Triangle{}},
		{PolyName: "Polygon 3", NumTriangles: 1, Triangles: []*pb.Triangle{}},
	}

	testPolygonList := []struct {
		polygons []*pb.Polygon
		err      error
		callErr  error
	}{
		{polygons: testPolygonSplice, err: nil, callErr: nil},
		{polygons: nil, err: errors.New("Recv"), callErr: nil},
		{polygons: nil, err: nil, callErr: errors.New("GetPolygonList")},
	}

	mockRendererClient := mock.NewMockRendererClient(ctrl)
	mockRenderer_ListOfPolygonsClient := mock.NewMockRenderer_ListOfPolygonsClient(ctrl)

	for _, testStruct := range testPolygonList {
		mockRendererClient.EXPECT().ListOfPolygons(gomock.Any(), nil).Return(mockRenderer_ListOfPolygonsClient, testStruct.callErr)
		for _, triangle := range testStruct.polygons {
			mockRenderer_ListOfPolygonsClient.EXPECT().Recv().Return(triangle, nil)
		}
		if testStruct.polygons == nil && testStruct.callErr == nil {
			mockRenderer_ListOfPolygonsClient.EXPECT().Recv().Return(nil, testStruct.err)
		} else if testStruct.callErr == nil {
			mockRenderer_ListOfPolygonsClient.EXPECT().Recv().Return(nil, io.EOF)
		}

		_, err := getPolygonList(mockRendererClient)
		if err != nil {
			log.Printf("Error caused by %v", err)
		}

	}

}

func TestClient_CreatePolygons(t *testing.T) {
	ctrl := gomock.NewController(t)

	testTriangleInputList := []*pb.Triangle{
		{TrId: 1, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 2, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 3, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 4, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		{TrId: 5, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		// {TrId: 6, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		// {TrId: 7, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		// {TrId: 8, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		// {TrId: 9, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		// {TrId: 10, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		// {TrId: 11, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
		// {TrId: 12, Vertice_1: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_2: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}, Vertice_3: &pb.Point{X: 100 * rand.Float32(), Y: 100 * rand.Float32()}},
	}

	testPolygonOutputList := []*pb.Polygon{
		{PolyName: "Polygon 1", NumTriangles: 2, Triangles: testTriangleInputList[0:1]},
		{PolyName: "Polygon 2", NumTriangles: 2, Triangles: testTriangleInputList[2:3]},
		{PolyName: "Polygon 3", NumTriangles: 1, Triangles: testTriangleInputList[4:]},
	}

	mockRendererClient := mock.NewMockRendererClient(ctrl)
	mockRenderer_CreatePoylgons := mock.NewMockRenderer_CreatePolygonsClient(ctrl)

	mockRendererClient.EXPECT().CreatePolygons(gomock.Any()).Return(mockRenderer_CreatePoylgons, nil)
	for index, triangle := range testTriangleInputList {
		mockRenderer_CreatePoylgons.EXPECT().Send(triangle).Return(nil)
		if (index+1)%2 == 0 {
			mockRenderer_CreatePoylgons.EXPECT().Recv().Return(testPolygonOutputList[(int)(index/2)], nil)
		}
	}
	mockRenderer_CreatePoylgons.EXPECT().Recv().Return(nil, io.EOF)
	mockRenderer_CreatePoylgons.EXPECT().CloseSend().Return(nil)

	err := executeCreatePolygon(mockRendererClient, testTriangleInputList)
	if err != nil {
		log.Printf("Error Caused by %v", err)
	}

}
