import { Link } from "react-router-dom";

function LogIn() {
  return (
    <div className="login">

      <h1>LogIn</h1>
      
      <form>
        <label htmlFor="email">Email</label>
        <input type="email" name="email" id="email" />
        <br/>
        
        <label htmlFor="password">Password</label>
        <input type="password" name="password" id="password" />
        <br/>

        <button type="submit">Log In</button>
        <br/>

        <button><Link to="/signup">Sign Up</Link></button>
        
      </form>
    </div>
  );
};

export default LogIn;