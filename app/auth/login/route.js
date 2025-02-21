export async function POST(req) {
    const { email, password } = await req.json();
  
    // Dummy user credentials
    const validUser = {
      email: "test@example.com",
      password: "password123",
    };
  
    if (email === validUser.email && password === validUser.password) {
      return Response.json({ message: "Login successful" }, { status: 200 });
    } else {
      return Response.json({ message: "Invalid credentials" }, { status: 401 });
    }
  }
  