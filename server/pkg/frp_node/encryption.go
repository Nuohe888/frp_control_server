package frp_node

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

// GenerateSign 根据指定的规则生成MD5签名。
// params: 待签名的参数map (key-value pairs)
// key: 商户密钥 (Secret Key)
// 返回: 小写格式的MD5签名字符串
// LINE: 14 函数定义
func GenerateSign(params map[string]string, secretKey string) string {
	// 1. 过滤并收集需要签名的Key
	var keys []string
	// LINE: 18 遍历参数
	for k, v := range params {
		// LINE: 20 过滤条件：排除 sign, sign_type 和空值
		if k == "sign" || k == "sign_type" || v == "" {
			continue
		}
		keys = append(keys, k)
	}

	// 2. 参数名ASCII码从小到大排序
	// LINE: 27 排序
	sort.Strings(keys)

	// 3. 拼接成URL键值对格式 (a=b&c=d)
	var builder strings.Builder
	// LINE: 32 遍历排序后的Key
	for i, k := range keys {
		value := params[k] // 从原始map获取value
		// LINE: 35 拼接键值对
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(value)
		// LINE: 39 如果不是最后一个参数，则追加 &
		if i < len(keys)-1 {
			builder.WriteString("&")
		}
	}

	// 4. 将商户密钥KEY拼接到字符串末尾
	// LINE: 46 拼接密钥
	builder.WriteString(secretKey)

	// 获取待签名字符串
	stringToSign := builder.String()

	// 调试时可以取消注释打印待签名字符串
	// fmt.Println("String to sign:", stringToSign)

	// 5. 计算MD5哈希值
	// LINE: 55 计算MD5
	hasher := md5.New()
	hasher.Write([]byte(stringToSign)) // Write接受[]byte
	hashBytes := hasher.Sum(nil)       // 获取哈希结果 []byte

	// 6. 转换为小写十六进制字符串
	// LINE: 60 转换为小写Hex字符串
	sign := hex.EncodeToString(hashBytes)

	return sign
}
