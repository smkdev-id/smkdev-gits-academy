import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Book from '../../Assets/Icons/Book.svg';
import OpenBook from '../../Assets/Icons/OpenBook.svg';
import ELearn from '../../Assets/Icons/ELearn.svg';

const Orientasi = () => {
  return (
    <Container className='orientasi-section-wrapper p-5'>
      <Row className='orientasi-section d-grid justify-content-center align-items-center'>
        <Col>
          <p className='text-center fs-3 fw-bold'>Orientasi Belajar SMKDEV</p>
          <p className='text-center fs-5'>
            Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat
            mengantarkan Anda menjadi talenta yang dibutuhkan oleh industri
            digital terkini.
          </p>
        </Col>
        <Col className='orientasi-image-container d-flex gap-3'>
          <Col>
            <Card xs={12} md={4} className='mb-4 border-0 shadow-sm h-100'>
              <Card.Img src={Book} class='p-2' width={100} alt='...' />
              <Card.Body>
                <Card.Title>Learning Path Industri</Card.Title>
                <Card.Text>
                  Learning path disusun berdasarkan kebutuhan dunia industri.
                </Card.Text>
              </Card.Body>
            </Card>
          </Col>
          <Col>
            <Card xs={12} md={4} className='mb-4 border-0 shadow-sm h-100'>
              <Card.Img src={OpenBook} class='p-2' width={100} alt='...' />
              <Card.Body>
                <Card.Title>Kurikulum Komprehensif</Card.Title>
                <Card.Text>
                  Kurikulum komprehensif dan senantiasa di perbarui berdasarkan
                  pengalaman di dunia industri
                </Card.Text>
              </Card.Body>
            </Card>
          </Col>
          <Col>
            <Card xs={12} md={4} className='mb-4 border-0 shadow-sm h-100'>
              <Card.Img src={ELearn} class='p-2' width={100} alt='...' />
              <Card.Body>
                <Card.Title>Project-Based Learning</Card.Title>
                <Card.Text>
                  Materi pembelajaran disusun dengan pendekatan - pendekatan
                  project based learning
                </Card.Text>
              </Card.Body>
            </Card>
          </Col>
        </Col>
      </Row>
    </Container>
  );
};

export default Orientasi;
