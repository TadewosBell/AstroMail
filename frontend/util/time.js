export function parseDateToJson(dateString) {
    // Parse the given date string
    const date = new Date(dateString);
  
    // Mapping for month and day abbreviations
    const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
    const days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
  
    // Extracting components
    const month = months[date.getMonth()];
    const day = days[date.getDay()];
    const dateFormatted = `${('0' + date.getDate()).slice(-2)}/${('0' + (date.getMonth() + 1)).slice(-2)}/${date.getFullYear()}`;
    const time = date.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', second: '2-digit' });
  
    // Constructing the JSON object
    const result = {
      month: month,
      day: day,
      date: dateFormatted,
      time: time,
      dateTime: date,
    };
  
    return result;
  }
  
  export function formatDateString(parsedDate) {
    // Create a date object from the parsed date
    const inputDate = parsedDate.dateTime;
  
    // Get the current date and time
    const currentDate = new Date();
  
    // Calculate the difference in milliseconds
    const diffMs = currentDate - inputDate;
  
    // Convert milliseconds to hours
    const diffHours = diffMs / (1000 * 60 * 60);
    // Check if the date is within the last 24 hours
    if (diffHours <= 12) {
      // Format and return the time part only if it's within the last 24 hours
      // Assuming the time in the parsedDate is already in local time format
      return new Date(inputDate).toLocaleTimeString('en-US', { hour: 'numeric', minute: '2-digit', hour12: true }).toLowerCase();
    } else {
      // Format and return the month and day if it's more than 24 hours ago
      return `${parsedDate.month} ${parseInt(parsedDate.date.split('/')[0], 10)}`;
    }
  }

  export function createDisplayDate(dateString){
    return formatDateString(parseDateToJson(dateString));
  }