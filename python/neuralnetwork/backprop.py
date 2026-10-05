x = 2.0
w = 3.0

target = 10.0

learning_rate = 0.1

# Forward Pass
prediction = w * x

loss = (prediction - target) ** 2

print("Prediction: ", prediction)
print("Loss: ", loss)

# Backpropagation
dl_dprediction = 2 * (prediction - target)  # derivate dl/dw
deprediction_dw = x
dL_dw = dl_dprediction * deprediction_dw
print("Gradient: ", dL_dw)

# Gradient descent
w = w - learning_rate * dL_dw
print("New weight: ", w)

# Forward Pass
prediction = w * x
loss = (prediction - target) ** 2
print("New Prediction: ", prediction)
print("New loss: ", loss)
