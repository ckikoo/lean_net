package sticky

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 随机生成字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestSt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000; i++ {
		// 生成随机的消息长度（1到100字节）
		messageLength := rand.Intn(1000000) + 1

		// 生成随机的消息内容
		testMessage := randomString(messageLength)

		// 使用缓冲区模拟网络传输
		buf := new(bytes.Buffer)

		// 创建编码器和解码器
		encoder := NewEncoder(buf)
		decoder := NewDecoder(buf)

		// 编码消息
		if err := encoder.Encode(testMessage); err != nil {
			t.Fatalf("Encoding failed: %v\n", err)
		}

		// 解码消息
		var decodedMessage string
		if err := decoder.Decode(&decodedMessage); err != nil {
			t.Fatalf("Decoding failed: %v\n", err)
		}

		// 验证结果
		if decodedMessage != testMessage {
			t.Fatalf("Test failed at iteration %d! Expected: %s, Got: %s\n", i, testMessage, decodedMessage)
		}
	}

	fmt.Println("All tests passed!")
}
