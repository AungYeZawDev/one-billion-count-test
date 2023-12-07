void main() {
  DateTime startTime = DateTime.now(); // Capture the start time
  for (int i = 0; i < 1000000000; i++) {
    print(i);
  }
  DateTime endTime = DateTime.now(); // Capture the end time
  Duration duration = endTime.difference(startTime); // Calculate the duration

  print('Time taken: ${duration.inSeconds} seconds');
}
