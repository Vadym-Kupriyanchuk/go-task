# go-task

### Input:
- A text file located on s server (https://siloh.pluto.tv/dev/task/input.txt), each line containing a single URL to a file(binary).
- Assume no authentication is needed for downloads.

### Functionality:

- Download all files concurrently, (_optimizing for memory and bandwidth usage_)
- Concatenate the contents of the downloaded files in the input order.
- Save the final output to output.mp4 on the local disk.
  
### Constraints:

- ___The number of URLs can be very high (e.g., 100k), so avoid memory overload.___
* ___Be efficient with IO and goroutines – no spawning 100k goroutines at once.___
- Ensure the download and concatenation preserve the order of URLs from the input list.
- Handle transient download failures with retries and timeouts.
- Log errors but continue processing the rest of the files.
