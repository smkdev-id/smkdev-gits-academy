import React from 'react';
import Logo from '../../Assets/Logo.png';
import { BsTwitter } from 'react-icons/bs';
import { SiLinkedin } from 'react-icons/si';
import { BsYoutube } from 'react-icons/bs';
import { FaFacebookF } from 'react-icons/fa';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Image from 'react-bootstrap/Image';

const Footer = () => {
  return (
    <Container className='footer-wrapper'>
      <Row>
        <Col className='footer-content'>
          <Col className='footer-logo-container'>
            <Image src={Logo} alt='' width={150} />
            <p className='fs-6 mt-1 fw-bold'>
              Creating High-Caliber Digital Talent
            </p>
          </Col>
          <Col className='footer-description'>
            <p className='fs-6'>
              Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang,
              Kec. Gedebage, Kota Bandung, Jawa Barat 40295
            </p>
          </Col>
          <Col className='footer-icons-group'>
            <p className='fw-bold fs-5'>Follow Us</p>
            <Col className='footer-icons d-flex gap-3'>
              <BsTwitter className='fs-3' />
              <SiLinkedin className='fs-3' />
              <BsYoutube className='fs-3' />
              <FaFacebookF className='fs-3' />
            </Col>
          </Col>
          <Col className='footer-copyright justify-content-center d-flex'>
            <p>
              &copy;2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights
              Reserved.
            </p>
          </Col>
        </Col>
      </Row>
    </Container>
  );
};

export default Footer;
