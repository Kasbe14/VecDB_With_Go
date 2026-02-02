Vector Database – High-Level Design

1. Purpose

This project implements a vector database whose responsibility is to:
Store vector embeddings efficiently
Search for vectors similar to a query vector
Return ranked results based on a similarity metric
The system is not responsible for data generation or transformation beyond vector search and retrieval.

2. Non-Goals

The system explicitly does not:
Generate new data
Perform generative AI tasks
Modify raw user data
Interpret semantic meaning of embeddings

3. Core Data Flow

Raw Data (text, image, audio, video, etc.)
   ↓
Embedder
   ↓
Vector (immutable)
   ↓
Index
   ↓
SearchResult (vector ID + similarity score)

4. Component Responsibilities

4.1 Embedder
The Embedder is responsible for:
Converting raw data into vector embeddings
Generating unique, stable vector IDs
Defining the embedding dimension
Defining the similarity metric used

Ownership Rule:
All vectors stored in the system MUST originate from an embedder.

4.2 Vector
A Vector is:
Immutable once created
Normalized at creation time
Associated with exactly one ID
Associated with exactly one dimension
Vectors expose read-only accessors only.

4.3 Index
Enforcing dimensional consistency
Performing similarity search
Returning ranked search results
The index does not:
Generate vectors
Modify vectors
Interpret raw data

5. Index Invariants
The index guarantees:
All stored vectors have the same dimension
The dimension is locked after the first insertion
Search results are sorted by similarity score (descending)
Search returns at most k results
The index does not guarantee:
At least one result
Non-empty results for empty indexes

6. Search Semantics
Inputs
Query vector (non-nil)
Integer k, where k > 0
Outputs
A slice of SearchResult
Each result contains:
Vector ID

Similarity score Behavior
If k > index size, return all vectors
If index is empty, return empty result set
If dimensions mismatch, return an error

7. Search vs Retrieval
Search identifies which vectors are relevant.
Retrieval (outside this system) uses vector IDs to fetch raw data.

8. Error Handling Philosophy
Invalid inputs fail fast with explicit errors
Valid inputs on empty data return empty results, not errors
Errors indicate contract violations, not absence of data

9. ID Ownership
Vector IDs are owned by the Embedder.
Users MUST NOT:
Manually create vectors
Manually assign vector IDs
Insert arbitrary vectors into an index

10. Current Scope (MVP)
Implemented:
Vector abstraction
Linear index
Add / Get / Delete
Dimension locking
Linear similarity search
Planned:
Embedder interface
Optimized search (top-k heap)
Metadata filtering
Persistent storage

🔒 Status
This document defines the current architectural contract.
Code MUST conform to this document.