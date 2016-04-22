package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var linuxConfig = ScanParsingConfig{linuxFindMac, linuxFindRssi}
var linuxConfigIwlist = ScanParsingConfig{linuxFindMac, linuxFindRssiIwList}

func TestLinuxOutEmptyResultWhenEmptyScan(t *testing.T) {
	data, _ := ParseOutput(linuxConfig, "")
	assert.True(t, len(data) == 0, "Result must be empty")
}

func TestLinuxOutSkipWhenInvalidMac(t *testing.T) {
	out := "BSS 801:37:73:ba:f7:dc\n" +
		"signal: dBm\n" +
		"BSS 11:11:11:aa:bb:cc\n" +
		"signal: -65.00 dBm"
	data, _ := ParseOutput(linuxConfig, out)

	expected := []WifiData{
		WifiData{"11:11:11:aa:bb:cc", -65},
	}

	assert.Equal(t, expected, data)
}

func TestLinuxOutSkipWhenInvalidSignal(t *testing.T) {
	out := "BSS 80:37:73:ba:f7:dc\n" +
		"signal: dBm\n" +
		"BSS 11:11:11:aa:bb:cc\n" +
		"signal: -65.00 dBm"
	data, _ := ParseOutput(linuxConfig, out)
	expected := []WifiData{
		WifiData{"11:11:11:aa:bb:cc", -65},
	}

	assert.Equal(t, expected, data)
}

func TestLinuxFullOutput(t *testing.T) {
	dat, _ := ioutil.ReadFile("test/linuxOutput.txt")
	data, err := ParseOutput(linuxConfig, string(dat))

	expected := []WifiData{
		WifiData{"80:37:73:ba:f7:dc", -25},
		WifiData{"80:37:73:87:46:82", -59},
		WifiData{"98:6b:3d:d7:84:e0", -60},
		WifiData{"08:95:2a:b1:e9:55", -76},
		WifiData{"2c:b0:5d:36:e3:b8", -54},
		WifiData{"58:20:b1:21:63:9f", -62},
		WifiData{"70:73:cb:bd:9f:b5", -78},
		WifiData{"b8:3e:59:78:35:99", -75},
		WifiData{"a0:63:91:2b:9e:64", -59},
		WifiData{"e0:3f:49:03:fd:38", -61},
		WifiData{"30:8d:99:71:95:c5", -78},
		WifiData{"80:37:73:ba:f7:d8", -37},
		WifiData{"a0:63:91:2b:9e:65", -55},
		WifiData{"80:37:73:87:56:36", -52},
		WifiData{"00:1a:1e:46:cd:11", -59},
		WifiData{"00:23:69:d4:47:9f", -59},
		WifiData{"54:65:de:6f:7e:d5", -70},
		WifiData{"70:73:cb:bd:9f:b6", -77},
		WifiData{"f8:35:dd:0a:da:be", -82},
		WifiData{"00:1a:1e:46:cd:10", -58},
		WifiData{"d4:05:98:57:b3:15", -79},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}

func TestPi3FullOutput(t *testing.T) {
	dat, _ := ioutil.ReadFile("test/pi3Output.txt")
	data, err := ParseOutput(linuxConfig, string(dat))

	expected := []WifiData{
		{"70:73:cb:bd:9f:b5", -72},
		{"4c:60:de:fe:e5:24", -80},
		{"80:37:73:ba:f7:d8", -16},
		{"a0:63:91:2b:9e:65", -43},
		{"00:23:69:d4:47:9f", -81},
		{"80:37:73:87:56:36", -68},
		{"2c:b0:5d:36:e3:b8", -75},
		{"58:20:b1:21:63:9f", -75},
		{"30:8d:99:71:95:c5", -81},
		{"c8:b3:73:25:22:51", -85},
		{"00:1a:1e:46:cd:10", -76},
		{"e0:46:9a:6d:02:ea", -91},
		{"08:95:2a:b1:e9:55", -81},
		{"00:1d:d4:7c:bd:30", -91},
		{"8c:09:f4:d3:84:50", -90},
		{"00:1a:1e:46:cd:11", -76},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}

func TestLinuxIwListOutput(t *testing.T) {
	dat, _ := ioutil.ReadFile("test/iwlistPiOutput.txt")
	data, err := ParseOutput(linuxConfigIwlist, string(dat))

	expected := []WifiData{

		{"70:73:cb:bd:9f:b5", -77},
		{"80:37:73:ba:f7:d8", -23},
		{"a0:63:91:2b:9e:65", -46},
		{"00:23:69:d4:47:9f", -71},
		{"80:37:73:87:56:36", -61},
		{"2c:b0:5d:36:e3:b8", -76},
		{"58:20:b1:21:63:9f", -69},
		{"30:8d:99:71:95:c5", -85},
		{"4c:60:de:fe:e5:24", -89},
		{"00:1a:1e:46:cd:10", -75},
		{"00:1a:1e:46:cd:11", -77},
		{"c0:c1:c0:f0:6f:cd", -89},
		{"c8:b3:73:25:22:51", -87},
		{"8c:09:f4:d3:84:50", -93},
		{"08:95:2a:b1:e9:55", -87},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, data)

}
