export function parseEmailAndName(input) {
    // Regular expression to match the name and email parts
    // This regex assumes the name, if present, is always before the email and wrapped in angle brackets
    const regex = /^(.*?)\s*(?:<([^>]+)>)?$/;
    const match = input.match(regex);
  
    // If there's no match, return nulls
    if (!match) {
      return { name: null, email: null };
    }
  
    // Extracting name and email from the match
    let name = match[1].trim() || null; // If the name part is empty, return null
    let email = match[2] || null; // If the email part is empty (unlikely in this context), return null
  
    // If the name is the same as the email, it means there was no name part in the input
    if (name === email) {
      name = null;
    }
  
    // Return an object with the name and email
    return { name, email };
}

export function previewString(str, numChars) {
    // Check if the string's length is greater than the specified number of characters
    if (str.length > numChars) {
      // Return the substring up to the specified number of characters plus "..."
      return str.substring(0, numChars) + "...";
    } else {
      // If the string is shorter or equal to the limit, return it as is
      return str;
    }
}