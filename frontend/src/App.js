import { Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Extra from "./pages/Extra";
import LogInPage from "./pages/LogInPage";
import SignUpPage from "./pages/SignUpPage";

export default function App() {
  return (
    <Routes>
      <Route path='/'>
        <Route index element={<Home />} />
        <Route path="login" element={<LogInPage />} />
        <Route path="signup" element={<SignUpPage />} />
        <Route path="extra" element={<Extra />} />
      </Route>
    </Routes>
  );
}
