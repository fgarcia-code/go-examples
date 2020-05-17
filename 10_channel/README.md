# Channels
## Concepts
* A channel provides a **mechanism** for **concurrently** executing **functions** to **communicate** by **sending** and **receiving** values of a specified element type
* Channels act as first-in-first-out queues
### Send statement
* A send statement **sends** a value **on a channel**
* **Communication blocks** until **the send can proceed** 
* **A send** on a **nil** channel **blocks forever**
### Receive operator
* A receive operator **receives** a value ***from a channel** 
* The **expression blocks** until a **value** is **available**
* **Receiving** from a **nil** channel **blocks forever**
### Close
* The built-in function **close(c)** records that **no more values** will be **sent** on the **channel** 
* It is an **error** if c is a **receive-only** channel. **Sending to** or **closing** a **closed channel** causes a **run-time panic**
* **Closing** the **nil** channel also causes a **run-time panic**
* After calling **close**, and after **any previously sent values have been received**, **receive operations** will **return** the **zero value** for the channel's type **without blocking**
## Simple Channels
A new, initialized channel value can be made using the **built-in** function **make**, which takes the **channel type** and an **optional capacity** as arguments.
## Buffer Channels
* The **capacity**, in number of elements, sets the **size** of the **buffer** in the channel
* If the **capacity** is **zero** or **absent**, the **channel** is **unbuffered** and **communication succeeds only when both a sender and receiver are ready**
* Otherwise, the **channel** is **buffered** and **communication succeeds without blocking** if **the buffer is not full (sends)** or **not empty (receives)**
* A **nil** channel is **never ready** for communication
## Directional Channels
The optional **<- operator** specifies the **channel direction**, **send** or **receive**. If **no direction** is given, the channel is **bidirectional**. 
## Range
* The **iteration values** produced are **the successive values sent** on the channel **until the channel is closed**
* If the **channel** is **nil**, the range expression **blocks forever**.
## Select
* A "select" statement **chooses** which of a **set** of **possible send or receive operations** will **proceed**
* **A case** may **assign** the result of **an expression** to **one** or **two variables**, which may be **declared** using a **short variable declaration**
* There can be **at most one default case** and it may appear anywhere in the list of cases.
* Since communication on **nil** channels can **never proceed**, a **select** with **only nil** channels and **no default** case **blocks forever**

**Execution:**
1. **All cases are evaluated exactly once**, in source order, **the result is a set of channels** to receive from or send to, and the corresponding values to send. Expressions with a short variable declaration or assignment are not yet evaluated.
2. If one or more of the communications can proceed, **a single one** that **can proceed** is **chosen via** a **uniform pseudo-random selection**. Otherwise, if there is a default case, that case is chosen. If there is **no default case**, **the "select" statement blocks** until **at least one** of the communications can **proceed**.
3. The respective communication operation is executed.
4. If the **selected case** is a **short variable declaration** or **an assignment**, **the expressions** are **evaluated** and **the received value** (or values) are **assigned**.
5. **The statement list** of the selected case **is executed**.
## ok Idiom
The value of **ok** is **true** if the **value received** was delivered by a **successful send** operation to the channel, or **false** if it is a **zero value** generated because the channel is ***closed** and **empty**.
## Fan-In
**A function** can read from **multiple inputs** and proceed **until** all are **closed** by **multiplexing** the **input channels** onto a **single channel** that's **closed** when **all the inputs are closed**. This is called fan-in.
## Fan-Out
**Multiple functions** can read from the **same channel** **until** that channel **is closed**; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.
## Context
Package context defines the **Context type**, which carries **deadlines**, **cancellation signals**, and other **request-scoped values** across **API boundaries** and **between processes**.