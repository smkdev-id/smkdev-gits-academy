import Col from "react-bootstrap/esm/Col";
import Container from "react-bootstrap/esm/Container";
import Row from 'react-bootstrap/Row';

const BenefitList = () => {
    return (
        <Container style={{ marginTop:"6vh" }}>
            <Row>
                <Col>
                    <h5>
                        💬 <br />
                        Bergabunglah dengan Komunitas
                    </h5>
                    <p>Terhubung dengan individu dan mentor yang berpikiran sama.</p>
                </Col>
                <Col>
                    <h5>
                        🏫 <br />
                        Belajar dari Ahli Industri
                    </h5>
                    <p>Dapatkan bimbingan dan bimbingan dari para profesional berpengalaman.</p>
                </Col>
                <Col>
                    <h5>
                        🌍 <br />
                        Selesaikan Proyek Dunia Nyata
                    </h5>
                    <p>Terapkan keahlian Anda untuk memecahkan tantangan TI praktis.</p>
                </Col>
            </Row>
        </Container>
    )
}

export default BenefitList;