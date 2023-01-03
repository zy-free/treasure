package iec

import "testing"

func TestByteCountSI(t *testing.T) {
    type args struct {
        b int64
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
        {
            name: "1",
            args: args{b: 900},
            want: "900B",
        },
        {
            name: "2",
            args: args{b: 2000},
            want: "2.00KB",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ByteCountSI(tt.args.b); got != tt.want {
                t.Errorf("ByteCountSI() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestByteCountIEC(t *testing.T) {
    type args struct {
        b int64
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
        {
            name: "1",
            args: args{b: 900},
            want: "900B",
        },
        {
            name: "2",
            args: args{b: 2000},
            want: "1.95KB",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ByteCountIEC(tt.args.b); got != tt.want {
                t.Errorf("ByteCountIEC() = %v, want %v", got, tt.want)
            }
        })
    }
}