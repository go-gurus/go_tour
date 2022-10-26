FROM webpronl/reveal-md:5.3.4

COPY resources /slides

CMD [ "/slides", "--theme", "theme/cc.css"]