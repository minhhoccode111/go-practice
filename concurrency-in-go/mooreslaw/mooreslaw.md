**Moore’s Law** refers to an observation made by Gordon Moore in 1965. It states that **the number of transistors on an integrated circuit doubles approximately every two years**, which historically led to **exponential growth in computing performance and reductions in cost per transistor**. For several decades, semiconductor manufacturers followed this trend by continuously shrinking transistor sizes using improved fabrication processes.

However, **Moore’s Law has slowed and is often considered no longer strictly true** because modern semiconductor technology has reached several **fundamental physical and economic limitations**.

---

# Physical Limitations That Stopped Moore’s Law

## 1️⃣ Atomic-Scale Limits

Transistors have been miniaturized to sizes **only a few nanometers wide**, which is approaching the scale of **individual atoms**.

When transistors become this small:

- There are **too few atoms** to reliably form stable structures.
- Manufacturing variations become significant.
- It becomes difficult to maintain predictable electrical behavior.

At this scale, further shrinking becomes physically impractical.

---

## 2️⃣ Quantum Tunneling

At extremely small transistor sizes, **quantum mechanical effects** appear.

Electrons can **tunnel through insulating barriers**, even when they are not supposed to conduct electricity.

Consequences:

- Current leaks through the transistor when it should be off.
- Power efficiency decreases.
- Circuits become unreliable.

This effect places a hard limit on how thin transistor gates can be.

---

## 3️⃣ Heat Dissipation (Power Density)

As more transistors are packed onto a chip:

- **Power density increases**
- Chips generate **more heat**

Removing heat becomes extremely difficult. Excess heat can:

- Damage components
- Reduce chip lifetime
- Force processors to throttle performance

This thermal limit prevented CPUs from continuing the previous trend of **increasing clock speeds**.

---

## 4️⃣ Signal Propagation Delays

Even if transistors could continue shrinking, **signal transmission across the chip becomes a bottleneck**.

Electrical signals must travel through metal interconnects, which have:

- Resistance
- Capacitance

As chips become denser, **communication between components slows relative to transistor switching speed**.

---

## 5️⃣ Manufacturing Complexity and Cost

Modern fabrication nodes (e.g., 5nm, 3nm) require extremely complex technologies such as:

- EUV lithography
- Multi-patterning
- Advanced materials

These technologies dramatically increase **fabrication cost and complexity**, making further scaling economically difficult.

---

# How the Industry Adapted

Since transistor scaling slowed, computing performance improvements now rely on:

- **Multi-core processors** (parallel execution)
- **Specialized processors** (GPUs, AI accelerators)
- **Improved architectures**
- **Better concurrency and parallel programming**

This shift toward **parallelism** is exactly why courses like _Concurrency in Go_ emphasize writing programs that can run efficiently on multiple cores. ⚙️💻

---

✅ **Summary**

Moore’s Law predicted that transistor density would double every two years, but it has slowed due to:

- Atomic-scale transistor limits
- Quantum tunneling effects
- Heat dissipation problems
- Signal propagation delays
- Increasing manufacturing complexity and cost

Because of these limits, the computing industry now focuses on **parallel computing and multi-core architectures rather than pure transistor scaling**.
