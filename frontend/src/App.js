import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Extra from "./pages/Extra";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<Home />} />
        <Route path="extra" element={<Extra />} />
      </Routes>
    </BrowserRouter>
  );
}
