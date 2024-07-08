import React from 'react';
import BannerImage from '../../Assets/hero-image.jpg';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

const Home = () => {
  return (
    <Container fluid className='home-container p-5'>
      <Row className='home-banner-container align-items-center'>
        <Col xs={12} md={6} className='home-text-section p-5 mt-5'>
          <p className='text-start fw-bold fs-1'>
            Jadilah Talenta Digital{' '}
            <span className='text-primary'>Masa Depan Indonesia</span>
          </p>
          <p className='text-start'>
            Belajar langsung dengan expert dari industri dengan kurikulum
            komprehensif berbasis project-based learning
          </p>
          <Button className='fw-bold'>Selengkapnya</Button>
        </Col>
        <Col xs={12} md={6} className='home-image-container p-5'>
          <img src={BannerImage} alt='Banner' className='img-fluid rounded-1' />
        </Col>
      </Row>
    </Container>
  );
};

export default Home;
