package main

func IsUp(v Flags) bool         { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)         { *v %^= FlagUp}
func SetBroadcast(v Flags) bool { *v |= FlagBroadcast }
func IsCast(v Flags) bool       { return v& (flagBroadcast|FlagMulticast) != 0}

func main() {
    _ = 1 << (10 * iota)
    KiB // 1024
    MiB // 1048576
    GiB // 1073741824
    TiB // 1099511627776
    PiB // 1125899906842624
    EiB // 1152921504606846976          (exceeds 1 << 32)
    ZiB // 1180591620717411303424       (exceeds 1 << 64)
    YiB // 1208925819614629174706176
}