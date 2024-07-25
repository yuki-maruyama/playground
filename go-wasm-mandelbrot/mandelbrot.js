function mandelbrot(c) {
  const iterations = 1000;
  const escapeRadius = 2.0;
  let z = { re: 0, im: 0 };

  for (let n = 0; n < iterations; n++) {
    const nextZ = {
      re: z.re * z.re - z.im * z.im + c.re,
      im: 2 * z.re * z.im + c.im
    };
    z = nextZ;
    if (Math.sqrt(z.re * z.re + z.im * z.im) > escapeRadius) {
      return n;
    }
  }
  return iterations;
}

function generateMandelbrot(width, height) {
  const xmin = -2.0, ymin = -1.5, xmax = 1.0, ymax = 1.5;
  let art = '';

  for (let py = 0; py < height; py++) {
    const y = ymin + (py / height) * (ymax - ymin);
    for (let px = 0; px < width; px++) {
      const x = xmin + (px / width) * (xmax - xmin);
      const c = { re: x, im: y };
      const m = mandelbrot(c);
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
      art += ch;
    }
    art += '\n';
  }
  return art;
}
