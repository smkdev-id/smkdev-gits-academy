import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import NavDropdown from 'react-bootstrap/NavDropdown';
import Button from 'react-bootstrap/Button';

const NavbarApps = () => {
    return (
        <>
            <Navbar expand="lg" className="bg-body-tertiary justify-content-between">
                <Container>
                    <div>
                        <Navbar.Brand href="/">
                            <img src="https://thumbs4.imagebam.com/a9/3a/e4/MEUKTN4_t.png" alt="logo (2).png"/>
                        </Navbar.Brand>
                    </div>
                    <div>
                        <Navbar.Toggle aria-controls="basic-navbar-nav" />
                        <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="me-auto">
                            <Nav.Link href="#home">About Us</Nav.Link>
                            <Nav.Link href="#link">Services</Nav.Link>
                            <Nav.Link href="#link">Projects</Nav.Link>
                            <NavDropdown title="More" id="basic-nav-dropdown">
                            <NavDropdown.Item href="#action/3.1">Bootcamp</NavDropdown.Item>
                            <NavDropdown.Item href="#action/3.2">Expert Class</NavDropdown.Item>
                            <NavDropdown.Item href="#action/3.3">Challenges</NavDropdown.Item>
                            </NavDropdown>
                            <Button variant="secondary mx-2">Sign In</Button>{' '}
                            <Button variant="outline-dark">Sign Up</Button>{' '}
                        </Nav>
                        </Navbar.Collapse>
                    </div>
                </Container>
            </Navbar>
        </>
    )
}

export default NavbarApps;