import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Image from 'react-bootstrap/Image';
import TalentAsset from '../../Assets/Talent.png';

const Talent = () => {
  return (
    <Container className='talent-section-wrapper p-5'>
      <Row className='talent-section d-grid justify-content-center align-items-center'>
        <Col>
          <p className='text-center fs-2 fw-bold'>Talenta Digital SMKDEV</p>
          <Image src={TalentAsset} width={750} />
        </Col>
      </Row>
    </Container>
  );
};

export default Talent;
