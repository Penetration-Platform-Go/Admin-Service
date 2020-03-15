package controller

import (
	"fmt"
	"strconv"
	"strings"

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

	wrongIP := checkEquipmentIPFormat(project.Equipment)
	wrongEquipmentPair := getConnectedEquipmentPairWithWrongIP(project.Equipment, network)

	lib.InitMatrix(network)
	deep := getNetworkDeep(network)
	equipmentTypeNumber := getPCNumber(project.Equipment)
	childNetwork := getChildNetworkNumber(network)

	information := ""

	if len(wrongIP) != 0 {
		information += fmt.Sprintf("%d Equipments' ip maybe wrong\n", len(wrongIP))
		for equipmentIndex := range wrongIP {
			information += fmt.Sprintf("%s: %s\n", project.Equipment[equipmentIndex].Name, wrongIP[equipmentIndex])
		}
	}

	wrongPairNumber := 0
	wrongPairInformation := ""
	for equipmentIndex := range wrongEquipmentPair {
		_, exists := wrongIP[equipmentIndex]
		if exists {
			continue
		}
		wrongPairNumber++
		wrongPairInformation += fmt.Sprintf("%s: %s <-> %s: %s\n",
			project.Equipment[equipmentIndex].Name,
			strings.Join(project.Equipment[equipmentIndex].Ip, ","),
			project.Equipment[wrongEquipmentPair[equipmentIndex]].Name,
			strings.Join(project.Equipment[wrongEquipmentPair[equipmentIndex]].Ip, ","))
	}
	if wrongPairNumber != 0 {
		information += fmt.Sprintf("\n%d pairs equipments are connected but those IP aren't in same network mask\n", wrongPairNumber)
		information += wrongPairInformation
	}

	information += fmt.Sprintf("\nNetwork Deep:\t %d\n", deep)
	information += "\nAll Type Equipment Number\n"
	for t := range equipmentTypeNumber {
		information += fmt.Sprintf("%s Number:\t %d\n", t, equipmentTypeNumber[t])
	}

	information += fmt.Sprintf("\nChild Network Number:\t %d\n\n", childNetwork)

	IPCorrectScore := 30 * len(wrongIP) / len(project.Equipment)
	CorrectConnectedScore := 20
	if len(wrongEquipmentPair) != 0 {
		CorrectConnectedScore = 0
	}
	EquipmentScore := equipmentTypeNumber["route"]*2 + equipmentTypeNumber["pc"] + equipmentTypeNumber["switch"]*3
	if EquipmentScore > 40 {
		EquipmentScore = 40
	}

	ChildNetworkScore := childNetwork * 3
	if ChildNetworkScore > 10 {
		ChildNetworkScore = 10
	}

	TotalScore := IPCorrectScore + CorrectConnectedScore + EquipmentScore + ChildNetworkScore
	return TotalScore, information
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

func checkEquipmentIPFormat(equipmentList []*pb.Equipment) map[int]string {
	var result map[int]string
	result = make(map[int]string)
	for index, each := range equipmentList {
		for _, ip := range each.Ip {
			if !checkIPFormat(ip) {
				result[index] = ip
			}
		}
	}
	return result
}

func checkIPFormat(ipString string) bool {
	var ipMap map[uint32]bool
	ipMap = map[uint32]bool{10: true, 100: true, 101: true, 172: true, 192: true, 127: true}
	ip, _ := getIPandMask(ipString)
	A := ip >> 24
	_, exists := ipMap[A]
	if !exists {
		return false
	}
	return true
}

func getConnectedEquipmentPairWithWrongIP(equipmentList []*pb.Equipment, network [][]bool) map[int]int {
	var result map[int]int
	result = make(map[int]int)

	count := len(network)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i == j {
				continue
			} else if network[i][j] {
				connected := checkIPConnected(equipmentList[i].Ip, equipmentList[j].Ip)
				if !connected {
					result[i] = j
				}
			}
		}
	}
	return result
}

func checkIPConnected(AIP []string, BIP []string) bool {
	// TODO: 先将 IP 初始化
	for i := 0; i < len(AIP); i++ {
		aip, amask := getIPandMask(AIP[i])
		for j := 0; j < len(BIP); j++ {
			bip, bmask := getIPandMask((BIP[j]))
			if aip>>uint((32-amask)) == bip>>uint((32-bmask)) {
				return true
			}
		}
	}
	return false
}

func getIPandMask(ip string) (uint32, int) {
	ipAndMask := strings.Split(ip, "/")
	ipList := strings.Split(ipAndMask[0], ".")
	mask := 24
	var result uint32
	result = 0
	if len(ipAndMask) == 2 {
		mask, _ = strconv.Atoi(ipAndMask[1])
	}
	for index, each := range ipList {
		temp, _ := strconv.Atoi(each)
		result += (uint32(temp) << uint(((3 - index) * 8)))
	}
	return result, mask
}
