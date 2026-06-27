package services

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// SubnetInfo 子网详细信息
type SubnetInfo struct {
	CIDR          string `json:"cidr"`
	Network       string `json:"network"`
	NetSize       int    `json:"netSize"`
	SubnetMask    string `json:"subnetMask"`
	WildcardMask  string `json:"wildcardMask"`
	Broadcast     string `json:"broadcast"`
	UsableStart   string `json:"usableStart"`
	UsableEnd     string `json:"usableEnd"`
	UsableCount   int64  `json:"usableCount"`
	HostCount     int64  `json:"hostCount"`
	BinaryNetwork string `json:"binaryNetwork"`
	BinaryMask    string `json:"binaryMask"`
}

// SubnetSplitResult 子网拆分结果
type SubnetSplitResult struct {
	Original SubnetInfo   `json:"original"`
	Subnets  []SubnetInfo `json:"subnets"`
}

func ipToInt(ip string) (int64, error) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return 0, fmt.Errorf("无效的 IP 地址: %s", ip)
	}
	var result int64
	for _, p := range parts {
		n, err := strconv.ParseInt(p, 10, 64)
		if err != nil || n < 0 || n > 255 {
			return 0, fmt.Errorf("无效的 IP 段: %s", p)
		}
		result = (result << 8) | n
	}
	return result, nil
}

func intToIP(n int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", (n>>24)&0xFF, (n>>16)&0xFF, (n>>8)&0xFF, n&0xFF)
}

// CalculateSubnet 计算 CIDR 子网的详细信息
func CalculateSubnet(cidr string) (SubnetInfo, error) {
	cidr = strings.TrimSpace(cidr)
	parts := strings.Split(cidr, "/")
	if len(parts) != 2 {
		return SubnetInfo{}, fmt.Errorf("无效的 CIDR 格式，请使用如 192.168.1.0/24")
	}

	ip := parts[0]
	netSize, err := strconv.Atoi(parts[1])
	if err != nil || netSize < 0 || netSize > 32 {
		return SubnetInfo{}, fmt.Errorf("无效的子网掩码位数: %s", parts[1])
	}

	ipInt, err := ipToInt(ip)
	if err != nil {
		return SubnetInfo{}, err
	}

	// 计算掩码
	var mask int64
	for i := 0; i < netSize; i++ {
		mask = (mask << 1) | 1
	}
	mask = mask << (32 - netSize)

	// 网络地址（对齐到网络边界）
	networkInt := ipInt & mask
	// 广播地址
	broadcastInt := networkInt | (int64(math.Pow(2, float64(32-netSize))) - 1)
	// 通配符掩码
	wildcardInt := ^mask & 0xFFFFFFFF

	hostCount := int64(math.Pow(2, float64(32-netSize)))
	usableCount := hostCount - 2

	networkIP := intToIP(networkInt)
	broadcastIP := intToIP(broadcastInt)
	maskIP := intToIP(mask)
	wildcardIP := intToIP(wildcardInt)

	var usableStart, usableEnd string
	if netSize >= 31 {
		usableStart = networkIP
		usableEnd = broadcastIP
		usableCount = hostCount
	} else {
		usableStart = intToIP(networkInt + 1)
		usableEnd = intToIP(broadcastInt - 1)
	}

	result := SubnetInfo{
		CIDR:          fmt.Sprintf("%s/%d", networkIP, netSize),
		Network:       networkIP,
		NetSize:       netSize,
		SubnetMask:    maskIP,
		WildcardMask:  wildcardIP,
		Broadcast:     broadcastIP,
		UsableStart:   usableStart,
		UsableEnd:     usableEnd,
		UsableCount:   usableCount,
		HostCount:     hostCount,
		BinaryNetwork: fmt.Sprintf("%032b", networkInt),
		BinaryMask:    fmt.Sprintf("%032b", mask),
	}
	return result, nil
}

// SplitSubnet 将 CIDR 拆分为指定子网掩码的多个子网
func SplitSubnet(cidr string, targetSize int) (SubnetSplitResult, error) {
	original, err := CalculateSubnet(cidr)
	if err != nil {
		return SubnetSplitResult{}, err
	}

	if targetSize <= original.NetSize {
		return SubnetSplitResult{}, fmt.Errorf("目标掩码 %d 必须大于原掩码 %d", targetSize, original.NetSize)
	}
	if targetSize > 32 {
		return SubnetSplitResult{}, fmt.Errorf("目标掩码不能超过 32")
	}

	ipInt, _ := ipToInt(original.Network)
	count := int(math.Pow(2, float64(targetSize-original.NetSize)))
	subnets := make([]SubnetInfo, 0, count)
	step := int64(math.Pow(2, float64(32-targetSize)))

	for i := 0; i < count; i++ {
		subIP := intToIP(ipInt + int64(i)*step)
		subCIDR := fmt.Sprintf("%s/%d", subIP, targetSize)
		info, err := CalculateSubnet(subCIDR)
		if err != nil {
			return SubnetSplitResult{}, err
		}
		subnets = append(subnets, info)
	}

	return SubnetSplitResult{
		Original: original,
		Subnets:  subnets,
	}, nil
}
