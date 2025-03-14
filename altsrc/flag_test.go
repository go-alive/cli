package altsrc

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/go-alive/cli"
)

type testApplyInputSource struct {
	Flag               FlagInputSourceExtension
	FlagName           string
	FlagSetName        string
	Expected           string
	ContextValueString string
	ContextValue       flag.Value
	EnvVarValue        string
	EnvVarName         string
	SourcePath         string
	MapValue           interface{}
}

func TestGenericApplyInputSourceValue(t *testing.T) {
	v := &Parser{"abc", "def"}
	c := runTest(t, testApplyInputSource{
		Flag:     NewGenericFlag(&cli.GenericFlag{Name: "test", Value: &Parser{}}),
		FlagName: "test",
		MapValue: v,
	})
	expect(t, v, c.Generic("test"))
}

func TestGenericApplyInputSourceMethodContextSet(t *testing.T) {
	p := &Parser{"abc", "def"}
	c := runTest(t, testApplyInputSource{
		Flag:               NewGenericFlag(&cli.GenericFlag{Name: "test", Value: &Parser{}}),
		FlagName:           "test",
		MapValue:           &Parser{"efg", "hig"},
		ContextValueString: p.String(),
	})
	expect(t, p, c.Generic("test"))
}

func TestGenericApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag: NewGenericFlag(&cli.GenericFlag{
			Name:    "test",
			Value:   &Parser{},
			EnvVars: []string{"TEST"},
		}),
		FlagName:    "test",
		MapValue:    &Parser{"efg", "hij"},
		EnvVarName:  "TEST",
		EnvVarValue: "abc,def",
	})
	expect(t, &Parser{"abc", "def"}, c.Generic("test"))
}

func TestStringSliceApplyInputSourceValue(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:     NewStringSliceFlag(&cli.StringSliceFlag{Name: "test"}),
		FlagName: "test",
		MapValue: []interface{}{"hello", "world"},
	})
	expect(t, c.StringSlice("test"), []string{"hello", "world"})
}

func TestStringSliceApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewStringSliceFlag(&cli.StringSliceFlag{Name: "test"}),
		FlagName:           "test",
		MapValue:           []interface{}{"hello", "world"},
		ContextValueString: "ohno",
	})
	expect(t, c.StringSlice("test"), []string{"ohno"})
}

func TestStringSliceApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewStringSliceFlag(&cli.StringSliceFlag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    []interface{}{"hello", "world"},
		EnvVarName:  "TEST",
		EnvVarValue: "oh,no",
	})
	expect(t, c.StringSlice("test"), []string{"oh", "no"})
}

func TestIntSliceApplyInputSourceValue(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:     NewIntSliceFlag(&cli.IntSliceFlag{Name: "test"}),
		FlagName: "test",
		MapValue: []interface{}{1, 2},
	})
	expect(t, c.IntSlice("test"), []int{1, 2})
}

func TestIntSliceApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewIntSliceFlag(&cli.IntSliceFlag{Name: "test"}),
		FlagName:           "test",
		MapValue:           []interface{}{1, 2},
		ContextValueString: "3",
	})
	expect(t, c.IntSlice("test"), []int{3})
}

func TestIntSliceApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewIntSliceFlag(&cli.IntSliceFlag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    []interface{}{1, 2},
		EnvVarName:  "TEST",
		EnvVarValue: "3,4",
	})
	expect(t, c.IntSlice("test"), []int{3, 4})
}

func TestBoolApplyInputSourceMethodSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:     NewBoolFlag(&cli.BoolFlag{Name: "test"}),
		FlagName: "test",
		MapValue: true,
	})
	expect(t, true, c.Bool("test"))
}

func TestBoolApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewBoolFlag(&cli.BoolFlag{Name: "test"}),
		FlagName:           "test",
		MapValue:           false,
		ContextValueString: "true",
	})
	expect(t, true, c.Bool("test"))
}

func TestBoolApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewBoolFlag(&cli.BoolFlag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    false,
		EnvVarName:  "TEST",
		EnvVarValue: "true",
	})
	expect(t, true, c.Bool("test"))
}

func TestStringApplyInputSourceMethodSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:     NewStringFlag(&cli.StringFlag{Name: "test"}),
		FlagName: "test",
		MapValue: "hello",
	})
	expect(t, "hello", c.String("test"))
}

func TestStringApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewStringFlag(&cli.StringFlag{Name: "test"}),
		FlagName:           "test",
		MapValue:           "hello",
		ContextValueString: "goodbye",
	})
	expect(t, "goodbye", c.String("test"))
}

func TestStringApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewStringFlag(&cli.StringFlag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    "hello",
		EnvVarName:  "TEST",
		EnvVarValue: "goodbye",
	})
	expect(t, "goodbye", c.String("test"))
}
func TestPathApplyInputSourceMethodSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:       NewPathFlag(&cli.PathFlag{Name: "test"}),
		FlagName:   "test",
		MapValue:   "hello",
		SourcePath: "/path/to/source/file",
	})

	expected := "/path/to/source/hello"
	if runtime.GOOS == "windows" {
		expected = `D:\path\to\source\hello`
	}
	expect(t, expected, c.String("test"))
}

func TestPathApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewPathFlag(&cli.PathFlag{Name: "test"}),
		FlagName:           "test",
		MapValue:           "hello",
		ContextValueString: "goodbye",
		SourcePath:         "/path/to/source/file",
	})
	expect(t, "goodbye", c.String("test"))
}

func TestPathApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewPathFlag(&cli.PathFlag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    "hello",
		EnvVarName:  "TEST",
		EnvVarValue: "goodbye",
		SourcePath:  "/path/to/source/file",
	})
	expect(t, "goodbye", c.String("test"))
}

func TestIntApplyInputSourceMethodSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:     NewIntFlag(&cli.IntFlag{Name: "test"}),
		FlagName: "test",
		MapValue: 15,
	})
	expect(t, 15, c.Int("test"))
}

func TestIntApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewIntFlag(&cli.IntFlag{Name: "test"}),
		FlagName:           "test",
		MapValue:           15,
		ContextValueString: "7",
	})
	expect(t, 7, c.Int("test"))
}

func TestIntApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewIntFlag(&cli.IntFlag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    15,
		EnvVarName:  "TEST",
		EnvVarValue: "12",
	})
	expect(t, 12, c.Int("test"))
}

func TestDurationApplyInputSourceMethodSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:     NewDurationFlag(&cli.DurationFlag{Name: "test"}),
		FlagName: "test",
		MapValue: 30 * time.Second,
	})
	expect(t, 30*time.Second, c.Duration("test"))
}

func TestDurationApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewDurationFlag(&cli.DurationFlag{Name: "test"}),
		FlagName:           "test",
		MapValue:           30 * time.Second,
		ContextValueString: (15 * time.Second).String(),
	})
	expect(t, 15*time.Second, c.Duration("test"))
}

func TestDurationApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewDurationFlag(&cli.DurationFlag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    30 * time.Second,
		EnvVarName:  "TEST",
		EnvVarValue: (15 * time.Second).String(),
	})
	expect(t, 15*time.Second, c.Duration("test"))
}

func TestFloat64ApplyInputSourceMethodSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:     NewFloat64Flag(&cli.Float64Flag{Name: "test"}),
		FlagName: "test",
		MapValue: 1.3,
	})
	expect(t, 1.3, c.Float64("test"))
}

func TestFloat64ApplyInputSourceMethodContextSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:               NewFloat64Flag(&cli.Float64Flag{Name: "test"}),
		FlagName:           "test",
		MapValue:           1.3,
		ContextValueString: fmt.Sprintf("%v", 1.4),
	})
	expect(t, 1.4, c.Float64("test"))
}

func TestFloat64ApplyInputSourceMethodEnvVarSet(t *testing.T) {
	c := runTest(t, testApplyInputSource{
		Flag:        NewFloat64Flag(&cli.Float64Flag{Name: "test", EnvVars: []string{"TEST"}}),
		FlagName:    "test",
		MapValue:    1.3,
		EnvVarName:  "TEST",
		EnvVarValue: fmt.Sprintf("%v", 1.4),
	})
	expect(t, 1.4, c.Float64("test"))
}

func runTest(t *testing.T, test testApplyInputSource) *cli.Context {
	inputSource := &MapInputSource{
		file:     test.SourcePath,
		valueMap: map[interface{}]interface{}{test.FlagName: test.MapValue},
	}
	set := flag.NewFlagSet(test.FlagSetName, flag.ContinueOnError)
	c := cli.NewContext(nil, set, nil)
	if test.EnvVarName != "" && test.EnvVarValue != "" {
		_ = os.Setenv(test.EnvVarName, test.EnvVarValue)
		defer os.Setenv(test.EnvVarName, "")
	}

	_ = test.Flag.Apply(set)
	if test.ContextValue != nil {
		f := set.Lookup(test.FlagName)
		f.Value = test.ContextValue
	}
	if test.ContextValueString != "" {
		_ = set.Set(test.FlagName, test.ContextValueString)
	}
	_ = test.Flag.ApplyInputSourceValue(c, inputSource)

	return c
}

type Parser [2]string

func (p *Parser) Set(value string) error {
	parts := strings.Split(value, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid format")
	}

	(*p)[0] = parts[0]
	(*p)[1] = parts[1]

	return nil
}

func (p *Parser) String() string {
	return fmt.Sprintf("%s,%s", p[0], p[1])
}
