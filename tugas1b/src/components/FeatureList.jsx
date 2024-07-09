import Col from "react-bootstrap/esm/Col";
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";

const FeatureList = () => {
    return (
        <Container style={{ marginTop:"8vh" }}>
            <Row style={{ justifyContent:"center", alignItems:"center" }}>
                <Col>
                    <h3>
                        Kembangkan Potensi Anda dengan Komunitas Kami
                    </h3>
                    <p>
                        Bergabunglah dengan komunitas kami dan dapatkan pengalaman proyek langsung, koneksi industri yang berharga, dan kembangkan keterampilan Anda untuk sukses di bidang TI.
                    </p>
                </Col>
                <Col>
                    <Row>
                        <Col>
                            <h4>
                                📝{""} <br />
                                Pengalaman Proyek Praktis
                            </h4>
                            <p>
                                Bekerja pada proyek IT dunia nyata dan dapatkan keterampilan praktis.
                            </p>
                        </Col>
                        <Col>
                            <h4>
                                🛜{""} <br />
                                Koneksi Industri
                            </h4>
                            <p>
                                Terhubung dengan para profesional dan perluas jaringan Anda.
                            </p>
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <h4>
                                🏫{""} <br />
                                Pengembangan Keterampilan
                            </h4>
                            <p>
                                Tingkatkan keterampilan TI Anda dan jadilah yang terdepan dalam industri.
                            </p>
                        </Col>
                        <Col>
                            <h4>
                                🪔{""} <br />
                                Temukan Potensi Anda dengan Pembelajaran Berbasis Proyek
                            </h4>
                            <p>
                                Bergabunglah dengan komunitas kami untuk mendapatkan pengalaman proyek langsung, terhubung dengan profesional industri, dan kembangkan keterampilan Anda.
                            </p>
                        </Col>
                    </Row>
                </Col>
            </Row>
        </Container>
    )
}

export default FeatureList;