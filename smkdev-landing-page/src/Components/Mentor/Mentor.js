import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Image from 'react-bootstrap/Image';
import Sudaryatno from '../../Assets/Mentor/Sudaryatno.jpeg';
import Helmi from '../../Assets/Mentor/Helmi.png';
import IbnuSina from '../../Assets/Mentor/IbnuSina.jpg';
import Rachmawati from '../../Assets/Mentor/Rachmawati.jpg';
import Samuel from '../../Assets/Mentor/Samuel.jpg';

const Testimoni = () => {
  return (
    <Container className='testimoni-section-wrapper p-5'>
      <Row className='testimoni-section d-grid justify-content-center align-items-center'>
        <Col className='testimoni-description'>
          <p className='text-center fs-2 fw-bold'>Expert Mentor</p>
          <p className='text-center fs-5'>
            SMKDEV menghadirkan expert dari industri untuk mendampingi proses
            belajarmu.
          </p>
        </Col>
        <Col className='testimoni-card d-flex gap-5 p-5'>
          <Col>
            <Card className='card w-100 border-0'>
              <Col>
                <Image src={IbnuSina} alt='...' className='rounded-circle' />
              </Col>
              <Card.Body class='card-body'>
                <h5 class='fs-6 fw-bold text-center'>Ibnu Sina Wardy</h5>
                <p class='fs-6 text-center fst-italic'>CEO </p>
                <div className='card-footer bg-transparent'>
                  <p class='fs-6 text-center fw-bold'>Gits Indonesia</p>
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col>
            <Card className='card w-100 border-0'>
              <Col>
                <Image src={Helmi} alt='...' className='rounded-circle' />
              </Col>
              <Card.Body class='card-body'>
                <h5 class='fs-6 fw-bold text-center'>Helmi Adi Prasetyo</h5>
                <p class='fs-6 text-center fst-italic'>Backend Developer</p>
                <div className='card-footer bg-transparent'>
                  <p class='fs-6 text-center fw-bold'>SMKDEV</p>
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col>
            <Card className='card w-100 border-0'>
              <Col>
                <Image src={Samuel} alt='...' className='rounded-circle' />
              </Col>
              <Card.Body class='card-body'>
                <h5 class='fs-6 fw-bold text-center'>
                  Samuel Pandohan Terampil Gultom
                </h5>
                <p class='fs-6 text-center fst-italic'>Backend Developer</p>
                <div className='card-footer bg-transparent'>
                  <p class='fs-6 text-center fw-bold'>SMKDEV</p>
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col>
            <Card className='card w-100 border-0'>
              <Col>
                <Image src={Rachmawati} alt='...' className='rounded-circle' />
              </Col>
              <Card.Body class='card-body'>
                <h5 class='fs-6 fw-bold text-center'>
                  Rachmawati Ari Taurisia
                </h5>
                <p class='fs-6 text-center fst-italic'>Curriculum Developer</p>
                <div className='card-footer bg-transparent'>
                  <p class='fs-6 text-center fw-bold'>SMKDEV</p>
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col>
            <Card className='card w-100 border-0'>
              <Col>
                <Image src={Sudaryatno} alt='...' className='rounded-circle' />
              </Col>
              <Card.Body class='card-body'>
                <h5 class='fs-6 fw-bold text-center'>Sudaryatno</h5>
                <p class='fs-6 text-center fst-italic'>CTO</p>
                <div className='card-footer bg-transparent'>
                  <p class='fs-6 text-center fw-bold'>GITS Indonesia</p>
                </div>
              </Card.Body>
            </Card>
          </Col>
        </Col>
      </Row>
    </Container>
  );
};

export default Testimoni;
