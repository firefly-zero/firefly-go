from pathlib import Path
from firefly_test import App, CLI

SDK_ROOT = Path(__file__).parent.parent
APP_ROOT = SDK_ROOT / "_examples" / "triangle"


def test_triangle(tmp_path: Path) -> None:
    cli = CLI(vfs=tmp_path)
    cli.build(APP_ROOT)
    app = App("demo.go-triangle")
    app.start()
    app.update()
    app.update()
