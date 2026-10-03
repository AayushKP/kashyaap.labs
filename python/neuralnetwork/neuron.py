import numpy as np


def relu(x):
    return max(0, x)


def neuron(inputs, weights, bias):
    weighted_sum = np.dot(inputs, weights) + bias
    return relu(weighted_sum)


inputs = np.array([4, -2, 5])
weights = np.array([0.1, 0.5, 0.2])

bias = -1

output = neuron(inputs, weights, bias)
print("Output: ", output)
