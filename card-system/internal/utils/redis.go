func InitRedis() error {
    client := redis.NewClient(&redis.Options{
        Addr: config.RedisAddr(),
    })
    _, err := client.Ping(context.Background()).Result()
    return err
}