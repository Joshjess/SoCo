import { Link } from "react-router-dom";

function NavBar() {
  return (
    <nav>
        <ul>
          <li>
            <Link to="/login">LogIn</Link>
          </li>
          <li>
            <Link to="/signup">SignUp</Link>
          </li>
        </ul>
      </nav>
  );
};

export default NavBar;