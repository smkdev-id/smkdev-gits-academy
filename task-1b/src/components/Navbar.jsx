import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Button from "react-bootstrap/Button";

function NavComponent() {
  return (
    <Navbar expand="lg" className="bg-body-transparant border-bottom">
      <Container>
        <Navbar.Brand href="#home">
          <img src="https://smkdev.storage.googleapis.com/wp/SMKDEV-Logo-Long-150x38.png" height={35} alt="" />
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="ms-auto gap-3 ">
            <Nav.Link href="#home" className="text-prim fw-semibold ">
              Learn
            </Nav.Link>
            <Nav.Link href="https://t.me/+xZ2kQTlLPMA1NjJl" target="_blank" className="text-prim fw-semibold">
              Community
            </Nav.Link>
            <Nav.Link href="https://www.smk.dev/blog/" target="_blank" className="text-prim fw-semibold">
              Blog
            </Nav.Link>
            <Nav.Link target="_blank" className="text-prim fw-semibold">
              FAQ
            </Nav.Link>
            <a href="https://lms.smk.dev/login" target="_blank" className="btn btn-primary px-4 bg-prim">
              Login
            </a>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default NavComponent;
