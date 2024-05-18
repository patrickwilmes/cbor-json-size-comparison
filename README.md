# CBOR JSON Size comparison

Just a tiny setup in go, to serialize the exact same structure once in cbor and json
and compare the resulting size. In the given example the resulting json is 138 bytes
long, where cbor is 120 bytes, so 18 bytes less. This might not seem as much
but accumulated to larger structures or bigger apis, this can actually make 
a difference.

# Resources
[Obsidian Note](obsidian://open?vault=second-brain&file=CBOR)
