package middleware

// 暂时没搞懂怎么用，先放在这
//func RecoveryMiddleware() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		defer func() {
//			if e := recover(); &e != nil {
//				response.Fail(ctx, nil, fmt.Sprint(e))
//			}
//		}()
//
//		ctx.Next()
//	}
//}
