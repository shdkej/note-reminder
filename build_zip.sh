mkdir -p build
pip install -t build -r recommender/requirments.txt
cd build
rm -rf *.dist-info __pycache__
find . -name "__pycache__" -type d -exec rm -rf {} +
rm -rf numpy* bin dateutil threadpoolctl.py six.py scipy*
cp ../event_handling/recommender_system.py .
zip -r ../event_handling/recommend.zip .
cd ../
