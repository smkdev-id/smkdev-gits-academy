import React from 'react';
import Logo from '../../Assets/Logo.png';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Button from 'react-bootstrap/Button';
import 'bootstrap/dist/css/bootstrap.min.css';

const NavbarComponents = () => {
  return (
    <>
      <Navbar
        expand='lg'
        sticky='top'
        className='p-3'
        style={{ backgroundColor: '#fcfcfc' }}
      >
        <Container>
          <Navbar.Brand href='#home'>
            <img
              src={Logo}
              width={150}
              className='d-inline-block align-top'
              alt='React Bootstrap logo'
            />
          </Navbar.Brand>
          <Navbar.Toggle aria-controls='basic-navbar-nav' />
          <Navbar.Collapse id='basic-navbar-nav'>
            <Nav className='ms-auto'>
              <Nav.Link href='' className='fw-bold'>
                Learn
              </Nav.Link>
              <Nav.Link href='' className='fw-bold'>
                Community
              </Nav.Link>
              <Nav.Link href='' className='fw-bold'>
                Blog
              </Nav.Link>
              <Button variant='primary' className='fw-bold'>
                Dashboard
              </Button>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </>
  );
};

export default NavbarComponents;
