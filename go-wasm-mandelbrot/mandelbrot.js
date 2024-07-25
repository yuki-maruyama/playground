function mandelbrot(c_re, c_im) {
  const iterations = 1000;
  const escapeRadius = 2.0;
  let z_re = 0, z_im = 0;
  let z_re_squared = 0, z_im_squared = 0;

  for (let n = 0; n < iterations; n++) {
    z_im = 2 * z_re * z_im + c_im;
    z_re = z_re_squared - z_im_squared + c_re;
    z_re_squared = z_re * z_re;
    z_im_squared = z_im * z_im;

    if (z_re_squared + z_im_squared > escapeRadius * escapeRadius) {
      return n;
    }
  }
  return iterations;
}

function generateMandelbrotJS(width, height) {
  const xmin = -2.0, ymin = -1.5, xmax = 1.0, ymax = 1.5;
  const dx = (xmax - xmin) / width;
  const dy = (ymax - ymin) / height;

  let art = [];

  for (let py = 0; py < height; py++) {
    const c_im = ymin + py * dy;
    let row = '';

    for (let px = 0; px < width; px++) {
      const c_re = xmin + px * dx;
      const m = mandelbrot(c_re, c_im);
      let ch;

      switch (true) {
        case m < 10:
          ch = ' ';
          break;
        case m < 20:
          ch = '.';
          break;
        case m < 40:
          ch = '*';
          break;
        case m < 100:
          ch = 'o';
          break;
        case m < 200:
          ch = 'O';
          break;
        default:
          ch = '@';
      }
      row += ch;
    }
    art.push(row);
  }
  return art.join('\n');
}

document.getElementById("runButton").addEventListener("click", async () => {
  const width = Number(document.getElementById("width").value);
  const height = Number(document.getElementById("height").value);
  const start = performance.now();
  const asciiArt = generateMandelbrotJS(width, height);
  const end = performance.now();
  document.getElementById("js_mandelbrot").textContent = asciiArt;
  document.getElementById("js_time").textContent = `JS Time: ${end - start}ms`;
});