package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Cahyo")
	}
}

func TestHelloWorldCahyo(t *testing.T) {

	result := HelloWorld("Cahyo")
	if result != "Hello Cahyo" {
		//Error
		t.Fatal("Result is not Hello Cahyo")
	}
}

func TestHelloWorldAssert(t *testing.T) {

	result := HelloWorld("Diha")
	assert.Equal(t, "Hello Diha", result, "Result must be Hello Diha")
}
func TestHelloWorldRequire(t *testing.T) {

	result := HelloWorld("Diha")
	require.Equal(t, "Hello Diha", result, "Result must be Hello Diha")
}

// func TestHelloWorldSkip(t *testing.T) {
// 	if runtime.GOOS == "darwin" {
// 		t.Skip("Can not run on mac os")
// 	}

// 	result := HelloWorld("Dihas")
// 	require.Equal(t, "Hello Diha", result, "Result must be Hello Diha")
// }

func TestSubTest(t *testing.T) {
	t.Run("Cahyo", func(t *testing.T) {
		result := HelloWorld("Cahyo")
		require.Equal(t, "Hello Cahyo", result, "Result must be Hello Cahyo")
	})
	t.Run("Diha", func(t *testing.T) {
		result := HelloWorld("Diha")
		require.Equal(t, "Hello Diha", result, "Result must be Hello Diha")
	})

}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Cahyo",
			Request:  "Cahyo",
			Expected: "Hello Cahyo",
		},
		{
			Name:     "Diha",
			Request:  "Diha",
			Expected: "Hello Diha",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Expected, result, fmt.Sprintf("Result must be %s", test.Expected))
		})
	}

}

func BenchmarkTableHelloWorld(b *testing.B) {
	bemchmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "Cahyo",
			Request: "Cahyo",
		},
		{
			Name:    "DihaDihaDihaDihaDihaDihaDihaDiha",
			Request: "DihaDihaDihaDihaDihaDihaDihaDiha",
		},
	}

	for _, test := range bemchmarks {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.Request)
			}
		})
	}

}
