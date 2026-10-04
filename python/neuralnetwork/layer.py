import numpy as np


def relu(x):
    return np.maximum(0, x)


# Single layer of neurons
inputs = np.array([2.0, 3.0, 1.0])

weights = np.array(
    [
        [0.5, 0.2, 0.8],  # Neuron 1
        [0.1, -0.4, 0.7],  # N2
        [0.3, 0.9, -0.2],  # N3
        [-0.5, 0.6, 0.4],  # N4
    ]
)

biases = np.array([0.1, 0.2, -0.3, 0.5])
# 4 neurons 4 outputs

z = np.dot(weights, inputs) + biases
print("Before activation:", z)

output = np.maximum(0, z)
print("After activation:", output)
