package main

import (
	"errors"
	"grpc/protobuf/pb"
	"io"
	"log"
	"net"
	"strconv"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var (
	ErrNoTriangleWithPoint = errors.New("No Triangle with a Point within")
)

type RendererServer struct {
	pb.UnimplementedRendererServer
	listTriangles []*pb.Triangle
	listPolygons  []*pb.Polygon
	triangleCount int
}

const trianglesPerPolygon = 2

func (r *RendererServer) SearchPoint(ctx context.Context, p *pb.Point) (*pb.Triangle, error) {
	for _, polygon := range r.listPolygons {
		for _, triangle := range polygon.Triangles {
			if proto.Equal(triangle.Vertice_1, p) ||
				proto.Equal(triangle.Vertice_2, p) ||
				proto.Equal(triangle.Vertice_3, p) {
				return triangle, nil
			}
		}
	}
	return nil, ErrNoTriangleWithPoint
}

func (r *RendererServer) CreatePolygons(stream pb.Renderer_CreatePolygonsServer) error {
	var trianglesOnPolygon []*pb.Triangle
	var polygonCounter int
	for {
		triangle, err := stream.Recv()
		if err == io.EOF {
			r.listPolygons = append(r.listPolygons, &pb.Polygon{
				PolyName:     "Polygon " + strconv.Itoa(polygonCounter),
				NumTriangles: int32(len(trianglesOnPolygon)),
				Triangles:    trianglesOnPolygon,
			})
			err := stream.Send(r.listPolygons[polygonCounter])
			if err != nil {
				return err
			}

			return nil
		}
		if err != nil {
			return err
		}
		trianglesOnPolygon = append(trianglesOnPolygon, triangle)
		if len(trianglesOnPolygon) == trianglesPerPolygon {
			r.listPolygons = append(r.listPolygons, &pb.Polygon{
				PolyName:     "Polygon " + strconv.Itoa(polygonCounter),
				NumTriangles: trianglesPerPolygon,
				Triangles:    append([]*pb.Triangle{}, trianglesOnPolygon...),
			})
			err := stream.Send(r.listPolygons[polygonCounter])
			if err != nil {
				return err
			}
			polygonCounter++
			trianglesOnPolygon = nil
		}
	}
}

func (r *RendererServer) CreateTriangle(stream pb.Renderer_CreateTriangleServer) error {
	var threePointsCounter int
	var triangleVertices []*pb.Point
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			if threePointsCounter < 3 {
				stream.SendAndClose(&pb.Triangle{
					TrId:      int32(r.triangleCount),
					Vertice_1: &pb.Point{},
					Vertice_2: &pb.Point{},
					Vertice_3: &pb.Point{},
				})
			}
			return err
		}

		if err != nil {
			return err
		}

		threePointsCounter++
		triangleVertices = append(triangleVertices, point)
		if threePointsCounter == 3 {
			threePointsCounter = 0
			r.listTriangles = append(r.listTriangles, &pb.Triangle{
				TrId:      int32(r.triangleCount),
				Vertice_1: triangleVertices[0],
				Vertice_2: triangleVertices[1],
				Vertice_3: triangleVertices[2],
			})
			err := stream.SendAndClose(r.listTriangles[r.triangleCount])
			r.triangleCount++
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func (r *RendererServer) GetPolyTriangles(p *pb.Polygon, stream pb.Renderer_GetPolyTrianglesServer) error {
	for _, polygon := range r.listPolygons {
		if proto.Equal(p, polygon) {
			for _, triangle := range polygon.Triangles {
				err := stream.Send(triangle)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return nil
}

func (r *RendererServer) ListOfTriangles(v *pb.Void, stream pb.Renderer_ListOfTrianglesServer) error {
	for _, triangle := range r.listTriangles {
		err := stream.Send(triangle)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RendererServer) ListOfPolygons(v *pb.Void, stream pb.Renderer_ListOfPolygonsServer) error {
	for _, polygon := range r.listPolygons {
		err := stream.Send(polygon)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Error has occoured %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRendererServer(grpcServer, &RendererServer{})
	grpcServer.Serve(lis)
}
