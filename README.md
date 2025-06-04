# 📘go-json-deepdive
Welcome to **go-json-deepdive** — a practical repository dedicated to exploring JSON in Go (Golang).  
It contains focused mini-projects and hands-on examples to help understand how to **parse**, **marshal**, **unmarshal**, **validate**, and **custom-handle** JSON data in Go.

## 🚀 What You'll Learn

- JSON encoding & decoding
- Custom MarshalJSON / UnmarshalJSON
- Working with nested structures
- Handling maps and slices
- Validation and pretty-printing
- Error handling best practices

## 📁 Project Structure
```
├── 01. Structured Data/    # Parse simple JSON objects into Go structs
├── 02. JSON Arrays/        # Handle JSON arrays and decode them into slices
├── 03. Nested Objects/     # Decode JSON with nested structures
├── 04. Primitive Types/    # Work with JSON numbers, strings, booleans
├── 05. Time Values/         # Decode date/time fields using time.Time
├── 06. Custom Parsing Logic/   # Add custom logic with UnmarshalJSON
├── 07. JSON Struct Tags - Custom Field Names/  # Struct tags to map JSON keys to Go fields
├── 08. Validating JSON Data/   # Check JSON format and handle validation errors
├── 09. Marshaling Structured Data/ # Convert Go structs into JSON
├── 10. Ignoring Empty Fields/  # Use omitempty to skip empty values in JSON
├── 11. Marshaling Slices/  # Encode slices/arrays into JSON
├── 12. Encoding Null Values/   # Represent nulls and pointer fields correctly
├── 13. Marshaling Maps/    # Convert Go maps into JSON objects
├── 14. Custom Encoding Logic/  # Customize JSON output using MarshalJSON
└── 15. Printing Formatted (Pretty-Printed) JSON/   # Pretty-print JSON with indentation for readability