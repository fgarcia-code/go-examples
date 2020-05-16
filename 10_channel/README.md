# Channels
## Simple Channels

## Buffer Channels

## Directional Channels

## Range

## Select

## ok Idiom

## Fan-In
**A function** can read from **multiple inputs** and proceed **until** all are **closed** by **multiplexing** the **input channels** onto a **single channel** that's **closed** when **all the inputs are closed**. This is called fan-in.
## Fan-Out
**Multiple functions** can read from the **same channel** **until** that channel **is closed**; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.
## Context
Package context defines the **Context type**, which carries **deadlines**, **cancellation signals**, and other **request-scoped values** across **API boundaries** and **between processes**.
