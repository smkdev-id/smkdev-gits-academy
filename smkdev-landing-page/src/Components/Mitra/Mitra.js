import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Image from 'react-bootstrap/Image';
import Eudeka from '../../Assets/Mitra/Eudeka.png';
import Digits from '../../Assets/Mitra/Digits.png';
import Arkana from '../../Assets/Mitra/Arkana.png';
import GitsIndonesia from '../../Assets/Mitra/GitsIndonesia.png';
import Hitopia from '../../Assets/Mitra/Hitopia.png';
import MantabOne from '../../Assets/Mitra/MantabOne.png';

const Mitra = () => {
  return (
    <Container className='justify-content-center align-item-center p-5'>
      <h1 className='text-center fs-3 fw-bold'>
        Dipercayai Oleh Mitra Industri
      </h1>
      <Row className='mitra-container d-flex justify-content-center mt-3'>
        <Col xs={6} md={4} xl={2} className='text-center'>
          <Image width={150} src={Digits} alt='' />
        </Col>
        <Col xs={6} md={4} xl={2} className='text-center'>
          <Image width={150} src={GitsIndonesia} alt='' />
        </Col>
        <Col xs={6} md={4} xl={2} className='text-center'>
          <Image width={150} src={Eudeka} alt='' />
        </Col>
        <Col xs={6} md={4} xl={2} className='text-center'>
          <Image width={150} src={Hitopia} alt='' />
        </Col>
        <Col xs={6} md={4} xl={2} className='text-center'>
          <Image width={150} src={Arkana} alt='' />
        </Col>
        <Col xs={6} md={4} xl={2} className='text-center'>
          <Image width={150} src={MantabOne} alt='' />
        </Col>
      </Row>
    </Container>
  );
};

export default Mitra;
