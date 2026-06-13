func maxProfit(prices []int) int {
    if len(prices) == 0 { return 0 }
    minPrice := math.MaxInt
    maxProfit := 0
    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        } else {
            newProfit := price - minPrice
            maxProfit = max(newProfit, maxProfit)
        }
    }
    return maxProfit
}
