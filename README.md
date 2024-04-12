## Job Queue

Currently uses `ring` as a queue to store Jobs before they are processed. It has O(1) time complexity for both push and pop operations. The disadvantage is the queue is fixed size. Can use `slice` or `list` if want to support unlimited queue size.

## Regenerating Code

To regenerate the code from `openapi.yaml`, run the following command:

```bash
$: make codegen
```
