package main

import (
	"context"
	"grpc/protobuf/pb"
	"io"
	"log"
	"math"
	"sort"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TriangleStruct struct {
	triangle *pb.Triangle
	side_a   float64
	side_b   float64
	side_c   float64
}

const (
	NonTriangleError = -1
)

func NewTriangleStruct(triangle *pb.Triangle) (*TriangleStruct, error) {
	ts := new(TriangleStruct)

	slice_sizes := []*float64{&ts.side_a, &ts.side_b, &ts.side_c}

	ts.side_a = math.Sqrt(math.Pow(float64(triangle.Vertice_2.X)-float64(triangle.Vertice_1.X), 2) + math.Pow(float64(triangle.Vertice_2.Y)-float64(triangle.Vertice_1.Y), 2))
	ts.side_b = math.Sqrt(math.Pow(float64(triangle.Vertice_3.X)-float64(triangle.Vertice_2.X), 2) + math.Pow(float64(triangle.Vertice_3.Y)-float64(triangle.Vertice_2.Y), 2))
	ts.side_c = math.Sqrt(math.Pow(float64(triangle.Vertice_1.X)-float64(triangle.Vertice_3.X), 2) + math.Pow(float64(triangle.Vertice_1.Y)-float64(triangle.Vertice_3.Y), 2))
	ts.triangle = triangle

	sort.Slice(slice_sizes, func(i, j int) bool {
		return *slice_sizes[i] > *slice_sizes[j]
	})

	if *slice_sizes[0] <= *slice_sizes[1]+*slice_sizes[2] {
		log.Println("It's a triangle!")
		return ts, nil
	}
	return ts, nil
}

func executeCreateTriangle(listPoints []*pb.Point, client pb.RendererClient) (*TriangleStruct, error) {
	stream, err := client.CreateTriangle(context.Background())
	if err != nil {
		log.Fatalf("CreateTriangle Failed %v", err)
	}
	log.Printf("Sending points from listOfPoints to CreateTriangle\n")
	for _, point := range listPoints {
		err := stream.Send(point)
		if err != nil {
			log.Fatalf("Error has occoured as points was being sent %v", err)
		}

	}
	triangle, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	triangleStructure, _ := NewTriangleStruct(triangle)

	log.Printf("Triangle created %+v, triangle itself: %+v\n", triangleStructure, triangleStructure.triangle)
	return triangleStructure, nil
}

func getPolygonList(client pb.RendererClient) ([]*pb.Polygon, error) {
	var polygonList []*pb.Polygon

	stream, err := client.ListOfPolygons(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	for {
		polygon, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		polygonList = append(polygonList, polygon)
		log.Printf("Received polygon %+v\n", polygon)
	}
	log.Println("Finished listing polygons!")
	return polygonList, nil
}

func getTriangleList(client pb.RendererClient) ([]*pb.Triangle, error) {
	var triangleList []*pb.Triangle
	stream, err := client.ListOfTriangles(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	for {
		triangle, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		triangleList = append(triangleList, triangle)
		log.Printf("Received triangle %+v\n", triangle)
	}

	log.Println("Finished listing triangles!")
	return triangleList, nil
}

func executeCreatePolygon(client pb.RendererClient, listTriangle []*pb.Triangle) error {
	var polygonList []*pb.Polygon
	waitc := make(chan struct{})

	stream, err := client.CreatePolygons(context.Background())
	if err != nil {
		return err
	}
	go func() {
		for {
			polygon, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error %v", err)
			}
			polygonList = append(polygonList, polygon)
			log.Printf("Polygon %+v received!\n", polygon)
		}
	}()

	for _, triangle := range listTriangle {
		err := stream.Send(triangle)
		if err != nil {
			return err
		}
		log.Printf("Triangle is sent %+v\n", triangle)
	}
	stream.CloseSend()
	<-waitc
	return nil
}
func executeGetPolyTriangles(client pb.RendererClient, in *pb.Polygon) error {
	stream, err := client.GetPolyTriangles(context.Background(), in)
	if err != nil {
		// log.Printf("Error %v", err)
		return err
	}
	for {
		triangle, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// log.Printf("Error on Recv %v", err)
			return err
		}
		log.Printf("Polygon %s has triangles %v\n", in.PolyName, triangle)
	}

	return nil
}

func main() {
	//	Starts connection with localhost on port 8081 utilizing dummy unsecure credentials
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Dial failed! %v", err)
	}
	defer conn.Close()

	//	Instantiates and initializes the client-side of the Renderer Service
	client := pb.NewRendererClient(conn)

	//	Creates 5 Triangles objects using mocked list of Points
	for counter := 0; counter < len(listOfPoints); counter += 3 {
		executeCreateTriangle(listOfPoints[counter:], client)
	}

	//	Gets a list of triangles currently registered on the server
	triangleList, err := getTriangleList(client)
	if err != nil {
		log.Fatalf("Error Occoured %v", triangleList)
	}

	//	Creates 2 Polygons(Each with 2 Triangles) using previously generated triangles
	err = executeCreatePolygon(client, triangleList)
	if err != nil {
		log.Fatalf("Error Occoured %v", err)
	}

	//	Gets a list of polygons currently registered on the server
	polygonList, err := getPolygonList(client)
	if err != nil {
		log.Fatalf("Error Occoured %v", err)
	}

	//	Listing all triangles in a specific polygon
	for _, polygon := range polygonList {
		err := executeGetPolyTriangles(client, polygon)
		if err != nil {
			log.Fatalf("Error Occoured %v", err)
		}
	}
	//
	// Searching for a unique Point inside all Polygons and Triangles, and returns the containing triangle
	triangle, err := client.SearchPoint(context.Background(), listOfPoints[1])
	if err != nil {
		log.Fatalf("Error Occoured %v", err)
	}
	log.Printf("Point %v is a vertice on %v triangle\n", listOfPoints[1], triangle)

}

var listOfPoints = []*pb.Point{
	{X: 2.4, Y: 1.3},
	{X: 5.2, Y: 3.1},
	{X: 8.5, Y: 9.2},
	{X: 75.3, Y: 23.1},
	{X: 86.4, Y: 44.5},
	{X: 12.3, Y: 67.4},
	{X: 62.4, Y: 23.86},
	{X: 34.4, Y: 18.35},
	{X: 96.4, Y: 78.8},
	{X: 23.4, Y: 37.5},
	{X: 77.4, Y: 64.8},
	{X: 59.4, Y: 89.1},
	{X: 46.4, Y: 76.2},
	{X: 23.4, Y: 54.6},
	{X: 67.4, Y: 34.87},
}
