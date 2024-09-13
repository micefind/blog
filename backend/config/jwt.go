package config

// JwtSecret 是用于签名和验证 JWT（JSON Web Token）的密钥
// 通常，密钥应该是高强度的随机值，以确保 JWT 的安全性
// 这里使用一个 64 字节的十六进制字符串作为密钥，转换为 byte 数组
var JwtSecret = []byte("6d59f7593b3e489da6a8de2d3e89f90b2d1c42e7243f4a2aa7b5aef1b3a07f92")
