# hasher

Use the 'HashPassword' function to generate a hashed value for the provided password.
h, err := hasher.HashPassword("password") // h == XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=, err == nil
Returns error if the provided password is empty string

Use the 'CheckPasswordHash' to verify if the provided password matches the provided hash.
matches := hasher.CheckPasswordHash("password", "XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=") // matches == true