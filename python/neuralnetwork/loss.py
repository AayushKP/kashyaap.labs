import numpy as np


# Mean Squared Error
def mse(predictions, targets):
    errors = predictions - targets
    squared_errors = errors**2

    return np.mean(squared_errors)


predictions = np.array([12.0, 18.0, 27.0])
targets = np.array([10.0, 20.0, 30.0])

loss = mse(predictions, targets)
print("Loss: ", loss)
