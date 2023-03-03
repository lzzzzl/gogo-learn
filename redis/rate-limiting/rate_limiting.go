package ratelimiting

// https://github.com/uptrace/bunrouter/blob/master/example/rate-limiting/main.go

// var (
// 	ctx            = context.Background()
// 	limiter        = redis_rate.NewLimiter(c.Client1())
// 	ErrRateLimited = errors.New("rate limited")
// )

// func rateLimit(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
// 	return func(w http.ResponseWriter, req bunrouter.Request) error {
// 		res, err := limiter.Allow(req.Context(), "project:123", redis_rate.PerMinute(10))
// 		if err != nil {
// 			return err
// 		}

// 		h := w.Header()
// 		h.Set("RateLimit-Remaining", strconv.Itoa(res.Remaining))

// 		if res.Allowed == 0 {
// 			// We are rate limited.

// 			seconds := int(res.RetryAfter / time.Second)
// 			h.Set("RateLimit-RetryAfter", strconv.Itoa(seconds))

// 			// Stop processing and return the error
// 			return ErrRateLimited
// 		}

// 		// Continue processing as normal.
// 		return next(w, req)
// 	}
// }

// func main() {
// 	res, err := limiter.Allow(ctx, "project:123", redis_rate.PerSecond(10))
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
// }
