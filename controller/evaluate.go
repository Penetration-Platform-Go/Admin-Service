package controller

import (
	"fmt"

	"github.com/Penetration-Platform-Go/Admin-Service/lib"
	"github.com/Penetration-Platform-Go/Admin-Service/model"
	pb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
)

func robotEvaluate(project *model.Project) (int, string) {
	mapTable := project.Map.Column
	var network [][]bool
	for _, each := range mapTable {
		network = append(network, each.Connected)
	}

	lib.InitMatrix(network)

	deep := getNetworkDeep(network)
	equipmentTypeNumber := getPCNumber(project.Equipment)
	childNetwrok := getChildNetworkNumber(network)

	information := fmt.Sprintf("Network Deep:\t %d\n", deep)
	for t := range equipmentTypeNumber {
		information += fmt.Sprintf("%s Number:\t %d\n", t, equipmentTypeNumber[t])
	}

	information += fmt.Sprintf("Child Network Number:\t %d\n", childNetwrok)

	return 0, information
}

func getChildNetworkNumber(network [][]bool) int {
	lib.Warshall(network)
	return lib.FindChildGraphNumber(network)
}

func getPCNumber(equipmentList []*pb.Equipment) map[string]int {
	var equipment map[string]int
	equipment = make(map[string]int)
	for _, each := range equipmentList {
		_, exists := equipment[each.Type]
		if exists {
			equipment[each.Type]++
		} else {
			equipment[each.Type] = 1
		}
	}
	return equipment
}

func getNetworkDeep(network [][]bool) int {
	distance := lib.Dijkstra(network)
	deep := 0
	for _, each := range distance {
		for _, dis := range each {
			if deep < dis && dis != 1000 {
				deep = dis
			}
		}
	}
	return deep
}
