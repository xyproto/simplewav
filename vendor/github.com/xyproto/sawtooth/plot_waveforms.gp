set title "Sawtooth Wave Comparison"
set xlabel "Sample"
set ylabel "Amplitude"
plot "formula_wave.dat" with lines title "Formula", "particle_wave.dat" with lines title "Particle"
