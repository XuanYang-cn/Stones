# Additive Quantization for Extreme Vector Compression

BY Artem Babenko and Victor Lempitsky CVPR 2014, the same author as PST.

## QUESTIONS

- What does extreme vector compression in the title mean? 极度有损压缩



## 1. MOTIVATION

- **New compression scheme:** approximating the vectors more accurately

## 2. IDEA

## 3. BACKGROUND

## 4. EXPERIMENT

## Discussion

Introduced new **lossy compression** scheme **Additive Quantization (AQ)** for high-dimension vectors. 

- **Accuracy:**  The approximation accuracy is higher (Figure 2), which translates into higher accuracy of ANNS and image classification applications.

- **Speed:** If evaluating scalar products, AQ is almost as fast as OPQ (Table 1) . Evaluating distances is slower due to the calculation of vectors length.
  $$
  ||x||^2
  $$

- **Compromise:** This length  is encoded using a separate byte. For example (Figure 5) they used 7 out of 8 bytes to store the AQ-compressed vector and the remaining 1 byte to store the length. 

  In this case, the accuracy will be sacrificed to increase the speed.

- **Drawbacks:** Vector encoding step is very complex and such improvement should be particularly considerable for the case of longer codes, for which they used hybrid algorithm combine AQ and PQ(APQ). The current hybrid scheme results int the fully connected graph, and they are currently investigating other approximations to the fully connected graphs.

## 5. CONCLUSION

