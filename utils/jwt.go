package util

var SecretKey = []byte("your-secret-key") // move to env/config

// func JWTMiddleware() echo.MiddlewareFunc {
// 	return middleware.JWTWithConfig(middleware.JWTConfig{
// 		SigningKey:    SecretKey,
// 		SigningMethod: "HS256",
// 	})
// }

// func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(jwt.MapClaims)

// 		if claims["role"] != "admin" {
// 			return echo.ErrUnauthorized
// 		}
// 		return next(c)
// 	}
// }
