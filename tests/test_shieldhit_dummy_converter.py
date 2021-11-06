import pytest
from os import path
from converter.shieldhit.parser import Parser
from converter.api import get_parser_from_str, run_parser

_Beam_str = """
RNDSEED      	89736501     ! Random seed
JPART0       	2            ! Incident particle type
TMAX0      	150.000000   0.0  ! Incident energy; (MeV/nucl)
NSTAT       1000    -1 ! NSTAT, Step of saving
STRAGG          2            ! Straggling: 0-Off 1-Gauss, 2-Vavilov
MSCAT           2            ! Mult. scatt 0-Off 1-Gauss, 2-Moliere
NUCRE           0            ! Nucl.Reac. switcher: 1-ON, 0-OFF
"""

_Mat_str = """MEDIUM 0
ICRU 276
END
"""

_Detect_str = """Geometry Cyl
    Name ScoringCylinder
    R 0.0 10.0 1
    Z 0.0 30.0 400
Geometry Mesh
    Name MyMesh_YZ
    X -5.0  5.0    1
    Y -5.0  5.0    100
    Z  0.0  30.0   300
Output
    Filename mesh.bdo
    Geo ScoringCylinder
    Quantity Dose
    """

_Geo_str = """
*---><---><--------><------------------------------------------------>
    0    0           protons, H2O 30 cm cylinder, r=10, 1 zone
*---><---><--------><--------><--------><--------><--------><-------->
  RCC    1       0.0       0.0       0.0       0.0       0.0      30.0
                10.0
  RCC    2       0.0       0.0      -5.0       0.0       0.0      35.0
                15.0
  RCC    3       0.0       0.0     -10.0       0.0       0.0      40.0
                20.0
  END
  001          +1
  002          +2     -1
  003          +3     -2
  END
* material codes: 1 - liquid water (ICRU material no 276), 1000 - vacuum, 0 - black body
    1    2    3
    1 1000    0
"""

_Test_dir = './test_runs'


@pytest.fixture
def parser() -> Parser:
    """Just a praser fixture."""
    return get_parser_from_str('dummy')


@pytest.fixture
def output_dir(tmp_path_factory, parser) -> str:
    """Fixture that creates a temporary dir for testing converter output and runs the conversion there."""
    output_dir = tmp_path_factory.mktemp(_Test_dir)
    run_parser(parser, {}, output_dir)
    return output_dir


def test_if_beam_created(output_dir) -> None:
    """Check if beam.dat file created"""
    with open(path.join(output_dir, 'beam.dat')) as f:
        assert f.read() == _Beam_str


def test_if_mat_created(output_dir) -> None:
    """Check if mat.dat file created"""
    with open(path.join(output_dir, 'mat.dat')) as f:
        assert f.read() == _Mat_str


def test_if_detect_created(output_dir) -> None:
    """Check if detect.dat file created"""
    with open(path.join(output_dir, 'detect.dat')) as f:
        assert f.read() == _Detect_str


def test_if_geo_created(output_dir) -> None:
    """Check if geo.dat file created"""
    with open(path.join(output_dir, 'geo.dat')) as f:
        assert f.read() == _Geo_str