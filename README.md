# VW Mushroom Go

## VW Contextual Bandit Simulator

This is a contextual bandit simulator implementing Vowpal Wabbit's RL framework through a Go interface.

(Under the hood this is using VW's command line interface)

For more details see Tutorials and Examples:

1. <https://github.com/VowpalWabbit/vowpal_wabbit>

2. <https://vowpalwabbit.org/>

### Data

The "mushroom" example uses the Mushroom dataset from UCI: <https://archive.ics.uci.edu/ml/datasets/Mushroom>

The "shuttle" example uses the StatLog dataset from UCI: <https://archive.ics.uci.edu/ml/datasets/Statlog+(Shuttle)>

The repo was originally designed for only the Mushroom dataset and then it was extended to allow for testing with any additional dataset.
