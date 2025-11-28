from pathlib import Path

import pytest
from firefly_test import CLI, App


SDK_ROOT = Path(__file__).parent.parent
EXAMPLES = list((SDK_ROOT / "_examples").iterdir())


@pytest.mark.parametrize("app_root", EXAMPLES)
def test_example(app_root: Path, tmp_path: Path) -> None:
    """For every example, test that it compiles and runs without panic.
    """
    cli = CLI(vfs=tmp_path)
    cli.build(app_root)
    app = App("demo.go-triangle")
    app.start()
    app.update()
