# Lab 5: Greedy Knapsack Algorithms

## Two Knapsack Problems:
1. **0-1 Knapsack**: Items cannot be split (take whole item or not)
2. **Fractional Knapsack**: Can take fractions of items

## Exercises:

### Exercise 1-6: Test Different Greedy Strategies for 0-1 Knapsack
Test these sorting orders (pick items in this order while weight ≤ capacity):
1. **Increasing weight** (lightest first)
2. **Decreasing value** (most valuable first)
3. **Decreasing (value - weight)**
4. **Decreasing value/weight ratio** (most value per kg first)

**Rough instructions for any strategy**:
1. Sort items by your chosen criteria
2. Start with empty knapsack
3. For each item in sorted order:
   - If adding it doesn't exceed capacity, add it
   - Otherwise, skip it

**Important**: For 0-1 Knapsack, greedy is NOT always optimal!
Example where greedy fails: Capacity=10, items: (value=60, weight=10), (value=100, weight=20)
- Greedy by value/weight takes first item: value=60
- Optimal takes second item: value=100 (but exceeds capacity? Actually this shows need for good example)

### Exercise 7-8: Small Cases
- Test with only 2 or 3 items
- Some greedy strategies might work for small cases

### Exercise 9: Capacity = 2
- When capacity is very small, some greedy strategies become optimal

### Exercise 10: Fractional Knapsack
**Rough instructions**:
1. Sort items by value/weight ratio (highest first)
2. For each item:
   - If whole item fits, take it all
   - Otherwise, take as much as fits (fraction)
   
**Important**: For fractional knapsack, greedy by value/weight ratio IS always optimal!